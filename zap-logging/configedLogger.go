package main

import (
	"encoding/json"

	"go.uber.org/zap"
)

func main() {
	jsonData := []byte(`{
    "level": "debug",
    "encoding": "json",
    "outputPaths": ["stdout", "log.txt"],
    "errorOutputPaths": ["stderr"],
    "initialFields": {"foo": "bar"},
    "encoderConfig": {
        "messageKey": "message",
        "levelKey": "level",
        "levelEncoder": "lowercase"
    }
    }`)

	configs := zap.Config{}
	if err := json.Unmarshal(jsonData, &configs); err != nil {
		panic(err)
	}

	logger, err := configs.Build()
	if err != nil {
		panic(err)
	}

	logger.Info("Configured logger")
	logger.Warn("Configured logger", zap.String("level", "info"))
}
