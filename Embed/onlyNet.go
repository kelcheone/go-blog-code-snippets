package main

import (
    _ "embed"
    "fmt"
 //   "io/fs"
    "log"
    "net/http"
)

//go:embed static/index.html
var indexHTML string

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, indexHTML)
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
