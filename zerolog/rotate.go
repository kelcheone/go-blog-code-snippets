package main

import (
    "gopkg.in/natefinch/lumberjack.v2"
    "github.com/rs/zerolog"
)

func main() {
    lumberjackLogger := &lumberjack.Logger{
        Filename: "./file.log",
        MaxSize: 1,
        MaxBackups: 3,
	    MaxAge:     28, // Days
	    Compress:   true,
    }

    logger := zerolog.New(lumberjackLogger).With().Timestamp().Logger()
    // of more than 1 MB, the log file will be rotated
    loop := 0

    for {
        logger.Info().Msg("I am logging to a file with rotation")
        if loop == 1000000 {
            break
        }
        loop++
    }
}

