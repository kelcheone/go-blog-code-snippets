package main

import (
    "errors"
    "os"

    "github.com/rs/zerolog"
)

func firstStack() error{
    err := secondStack()
    return err
}

func stackTrace() error {
    return errors.New("Something went wrong")
}

func main() {
    err := firstStack()

    logger := zerolog.New(os.Stdout)
    logger.Error().Err(err).Stack().Msg("")
}o
