package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

const APIVersion = "0.4.0"

func configReader(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	return nil
}

func readConfig(n int, logger *zap.Logger) {
	for i := 0; i < n; i++ {
		filename := fmt.Sprintf("config-file-%d.cfg", i)
		if err := configReader(filename); err != nil {
			logger.Error(
				fmt.Sprintf("error while reading config from %s", filename),
				zap.String("filename", filename),
				zap.Error(err),
			)
		}
	}
}

func main() {
	logger, _ := zap.NewProduction()

	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	logger.Info("ova-promise-api", zap.String("version", APIVersion))

	readConfig(1, logger)
}
