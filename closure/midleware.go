package main

import (
    "fmt"
    "net/http"
)

func midleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Before")
        next.ServeHTTP(w, r)
        fmt.Println("After")
    })
}

func main() {
    http.Handle("/", midleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Hello World")
    })))
    http.ListenAndServe(":8080", nil)
}
