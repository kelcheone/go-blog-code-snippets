package main

import (
	"log"

	"go.uber.org/zap"
)

func id() {

	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatalf("Could not instanciate zap %v", err)
	}

	defer logger.Sync()

	//logger.Info("Hello World!")
	//logger.Info("Go can be fun!", zap.String("level", "info"))

	//logger.Info("Go can be fun!", zap.String("level", "info"),
	//zap.Int("count", 1), zap.Bool("is_active", true), zap.Float64("price",10.99))

	log.Print("-------------------------------------------------------------------------------------------")

	//logger.Debug("Go can be fun!")

	//logger.Info("Go can be fun!")

	//logger.Warn("Go can be fun!")

	//logger.Error("Go can be fun!")

	//logger.DPanic("Go can be fun!")

	logger.Panic("Go can be fun!")

	// logger.Fatal("Go can be fun!")
}
