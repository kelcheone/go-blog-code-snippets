package main

import (
    "log"
    "net/http"
)


// We don't know the request id yet, so we will use a placeholder
func loggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Request id: %s", "request-id")
        next.ServeHTTP(w, r)
    })
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Hello World")
    })

    http.ListenAndServe(":8080", loggerMiddleware(http.DefaultServeMux))
}
