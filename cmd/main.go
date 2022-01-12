package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/rabdavinci/fibo/data"
	pb "github.com/rabdavinci/fibo/gen/proto"
	"github.com/rabdavinci/fibo/handlers"
	"google.golang.org/grpc"
)

func main() {
	l := log.New(os.Stdout, "fibo-api ", log.LstdFlags)
	fc := data.FiboCache{}

	lh := handlers.NewFibo(l, fc)

	sm := mux.NewRouter()

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", lh.CalculateSlice)
	postRouter.Use(lh.MiddlewareValidate)

	// create a new server
	s := http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting http server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// create grpc server
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		l.Fatalln(err)
	}
	grpcserver := grpc.NewServer()

	pb.RegisterFiboApiServer(grpcserver, &handlers.FiboApiServer{})

	go func() {
		l.Println("Starting grpc server on port 8080")
		err = grpcserver.Serve(listener)
		if err != nil {
			l.Fatalln(err)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
