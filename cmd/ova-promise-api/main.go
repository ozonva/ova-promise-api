package main

import (
	"go.uber.org/zap"
)

const ApiVersion = "0.0.1"

func main() {
	logger, _ := zap.NewProduction()

	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	logger.Info("ova-promise-api", zap.String("version", ApiVersion))
}
