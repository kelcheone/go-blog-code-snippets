package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Define a custom level for the "Notice" level
	noticeLevel := zap.NewAtomicLevelAt(zapcore.InfoLevel + 1)

	// Create a new logger with a custom "Notice" level
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(os.Stdout),
		noticeLevel,
	))

	// Log a message with the custom "Notice" level
	logger.With(zap.String("level", "notice")).Notice("This is a custom notice message")

	// Log a message with the standard "Info" level
	logger.Info("This is a standard info message")

	logger.Sync()
}
