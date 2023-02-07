package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	// "go.uber.org/zap/zapcore"
)

func main() {
	atom := zap.NewAtomicLevel()

	atom.SetLevel(zapcore.ErrorLevel)

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(os.Stdout),
		atom,
	))

	logger.Info("Hello World!")

}
