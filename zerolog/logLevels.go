package main

import (
    "os"

    "github.com/rs/zerolog"
)

func main() {
    logger := zerolog.New(os.Stdout)
    logger.Info().Msg("Hello World!")
    logger.Debug().Msg("Hello World!")
    logger.Error().Msg("Hello World!")
    logger.Fatal().Msg("Hello World!")
    logger.Panic().Msg("Hello World!")
}
