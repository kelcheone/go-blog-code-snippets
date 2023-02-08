package main

import (
    _ "embed"
    "fmt"
    "log"
    "net/http"
)

//go:embed static/index.html
var indexHTML string

//go:embed static/robots.txt
var robotsTXT string

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, indexHTML)
    })

    http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, robotsTXT)
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
