package main

import (
    "os"

    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func main() {
    log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

    log.Info().Msg("Hello World!")
}
