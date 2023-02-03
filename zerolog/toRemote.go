package main

import (
    "net"
    "os"

    "github.com/rs/zerolog"
)

func main() {
    con, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        panic(err)
    }

    defer con.Close()

    logger := zerolog.New(con).With().Timestamp().Logger()

    logger.Info().Msg("I am logging to a remote server")

    os.Exit(0)
}
