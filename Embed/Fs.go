package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexHTML, err := staticFiles.ReadFile("static/index.html")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, string(indexHTML))
	})

	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		robotsTXT, err := staticFiles.ReadFile("static/robots.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, string(robotsTXT))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
