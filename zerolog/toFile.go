package main

import (
    "os"

    "github.com/rs/zerolog"
)

func main() {
    file, err := os.OpenFile("file.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        panic(err)
    }

    defer file.Close()

    logger := zerolog.New(file).With().Timestamp().Logger()
    logger.Info().Msg("I am logging to a file")
}
