package main

import (
    "os"

    "github.com/rs/zerolog"
)

func main() {
    zerolog.SetGlobalLevel(zerolog.ErrorLevel)
    logger := zerolog.New(os.Stdout)
    logger.Info().Msg("This is an info message")
    logger.Debug().Msg("This is a debug message")
    logger.Error().Msg("This is an error message")
    logger.Warn().Msg("This is a warn message")
}
