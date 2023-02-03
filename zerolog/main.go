package main

import (
    "os"

    "github.com/rs/zerolog"
)

func main() {
    logger := zerolog.New(os.Stdout)
    logger.Info().Msg("Hello World!")
}
