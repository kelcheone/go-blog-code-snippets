package main

import (
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	suggeredLogger := logger.Sugar()

	suggeredLogger.Infow("Failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "https://example.com",
		"attempt", 3,
		"backoff", 1,
	)

	suggeredLogger.Infof("Failed to fetch URL: %s", "https://example.com")
}
