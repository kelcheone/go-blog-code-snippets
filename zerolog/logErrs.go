package main

import (
    "errors"
    "os"

    "github.com/rs/zerolog"
)

func main() {
    err := errors.New("Something went wrong")

    logger := zerolog.New(os.Stdout)
    logger.Error().Err(err).Msg("")
}
