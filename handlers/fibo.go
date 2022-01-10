package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/rabdavinci/fibo/data"
)

type Fibo struct {
	l *log.Logger
}

func NewFibo(l *log.Logger) *Fibo {
	return &Fibo{l}
}

func (f *Fibo) CalculateSlice(rw http.ResponseWriter, r *http.Request) {
	f.l.Println("Handle POST")

	fi := r.Context().Value(KeyFibo{}).(data.FiboInterval)
	fl := data.CalculateSlice(&fi)

	err := fl.ToJSON(rw)
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
