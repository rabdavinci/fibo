package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/rabdavinci/fibo/data"
)

type Fibo struct {
	l  *log.Logger
	fc *redis.Client
}

func NewFibo(l *log.Logger, fc *redis.Client) *Fibo {
	return &Fibo{l, fc}
}

func (f *Fibo) CalculateSlice(rw http.ResponseWriter, r *http.Request) {
	f.l.Println("Handle POST")
	fi := r.Context().Value(KeyFibo{}).(data.FiboInterval)
	var fl = data.FiboList{}
	key := fmt.Sprintf("%v:%v", fi.X, fi.Y)
	cv, err := f.fc.Get(key).Result()
	if err != nil && err != redis.Nil {
		f.l.Fatalln(err)
	} else if err == redis.Nil {
		fl.List = data.CalculateSlice(&fi)
		err = f.fc.Set(key, fl.List, 0).Err()
		if err != nil {
			f.l.Fatalln(err)
		}
		f.l.Printf("Generated for x=%v y=%v value%v", fi.X, fi.Y, fl)
	} else {
		fl.List = cv
		f.l.Printf("Got from Cache for x=%v y=%v value%v", fi.X, fi.Y, fl)
	}

	err = fl.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

type KeyFibo struct{}

func (f Fibo) MiddlewareValidate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fi := data.FiboInterval{}
		err := fi.FromJSON(r.Body)
		if err != nil {
			f.l.Println("[ERROR] deserializing params", err)
			http.Error(rw, "Error reading params", http.StatusBadRequest)
			return
		}

		// add the fibo interval to the context
		ctx := context.WithValue(r.Context(), KeyFibo{}, fi)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
