package main

import "go.uber.org/zap"

const APIVersion = "0.2.0"

func main() {
	logger, _ := zap.NewProduction()

	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	logger.Info("ova-promise-api", zap.String("version", APIVersion))
}
