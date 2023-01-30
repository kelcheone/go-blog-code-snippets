package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func enrichContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled")
		err := ctx.Err()
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	case <-time.After(5 * time.Second):
		fmt.Println("Context not cancelled")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", enrichContext(http.DefaultServeMux))
}
