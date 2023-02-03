package main

import (
    "net/http"

    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func main() {
    log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })

    http.ListenAndServe(":8080", zerolog.Handler(mux))
}
