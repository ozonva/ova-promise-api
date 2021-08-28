package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"

	promiseRepo "github.com/ozonva/ova-promise-api/internal/implementation/pg.repository"
	"github.com/ozonva/ova-promise-api/internal/infrastructure"
	"github.com/ozonva/ova-promise-api/internal/usecase"
)

const APIVersion = "0.7.0"

//nolint //task
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

//nolint //task
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

const (
	chunkSize = 10
	address   = "127.0.0.1:9001"
)

func main() {
	logger, _ := zap.NewProduction()

	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	logger.Info("ova-promise-api", zap.String("version", APIVersion))

	pgConString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
	)

	config, err := pgxpool.ParseConfig(pgConString)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to parse DATABASE_URL: %v\n", err)
		panic(err)
	}

	config.MaxConns = 10
	config.ConnConfig.PreferSimpleProtocol = true

	dbPool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	}
	defer dbPool.Close()

	ucHandler := usecase.HandlerConstructor{
		PromiseRepository: promiseRepo.CreateRepository(dbPool),
		ChunkSize:         chunkSize,
		Logger:            logger,
	}.New()

	server := infrastructure.InitGRPCServer(ucHandler, logger)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("starting grpc server")

	if err := server.Serve(listener); err != nil {
		logger.Fatal(err.Error())
	}
}
