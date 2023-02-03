package main

import (
    "os"

    "github.com/rs/zerolog"
)

func main() {
    logger := zerolog.New(os.Stdout)
    logger.Level(zerolog.InfoLevel)
    logger.Info().Msg("Hello World!")
    logger.Debug().Msg("Hello World!")
    logger.Error().Msg("Hello World!")
}
