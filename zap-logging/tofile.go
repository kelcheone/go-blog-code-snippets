package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	logger := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.Lock(file),
			zapcore.DebugLevel,
		),
	)

	logger.Info("Hello World!")
	logger.Info("Go can be fun!", zap.String("level", "info"))
	logger.Info("Go can be fun!", zap.String("level", "info"),
		zap.Int("count", 1), zap.Bool("is_active", true), zap.Float64("price", 10.99))

}
