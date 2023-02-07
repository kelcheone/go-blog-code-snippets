package main

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Could not open file %v", err)
	}

	defer file.Close()

	logger := zap.New(
		zapcore.NewTee(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
				zapcore.Lock(file),
				zapcore.DebugLevel,
			),
			zapcore.NewCore(
				zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
				zapcore.Lock(os.Stdout),
				zapcore.DebugLevel,
			),
		),
	)

	logger.Info("We are logging to file and stdout")
	logger.Info("We are logging to file and stdout", zap.String("level", "info"))
	logger.Info("We are logging to file and stdout", zap.String("level", "info"),
		zap.Int("count", 1), zap.Bool("is_active", true), zap.Float64("price", 10.99))
}
