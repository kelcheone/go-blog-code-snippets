package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	atomic := zap.NewAtomicLevel()

	atomic.SetLevel(zapcore.DebugLevel)

	enconfig := zap.NewProductionEncoderConfig()
	enconfig.EncodeTime = zapcore.ISO8601TimeEncoder
	enconfig.TimeKey = "timestamp"

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(enconfig),
		zapcore.Lock(os.Stdout),
		atomic,
	))

	logger.Debug("Change time format")
	logger.Info("Change time format")

	logger.Sync()

}
