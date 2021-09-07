package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	eventProducer "github.com/ozonva/ova-promise-api/internal/implementation/kafka.producer"
	prometheusmetrics "github.com/ozonva/ova-promise-api/internal/implementation/prometheus.metrics"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"

	promiseRepo "github.com/ozonva/ova-promise-api/internal/implementation/pg.repository"
	"github.com/ozonva/ova-promise-api/internal/infrastructure"
	"github.com/ozonva/ova-promise-api/internal/usecase"
)

const APIVersion = "0.8.0"

//nolint // deprecated
func configReader(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	return nil
}

//nolint // deprecated
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

//nolint // main func may be dirty
func main() {
	ctx := context.Background()

	logger, _ := zap.NewProduction()

	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	logger.Info("ova-promise-api", zap.String("version", APIVersion))

	cfg, err := infrastructure.NewConfigFromEnv()
	if err != nil {
		logger.Fatal(err.Error())
	}

	dbConfig, err := pgxpool.ParseConfig(cfg.Database.ConnectionString)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to parse DATABASE_URL: %v\n", err)
		logger.Fatal(err.Error())
	}

	dbConfig.MaxConns = cfg.Database.MaxConn
	dbConfig.ConnConfig.PreferSimpleProtocol = true

	dbPool, err := pgxpool.ConnectConfig(ctx, dbConfig)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		logger.Fatal(err.Error())
	}
	defer dbPool.Close()

	kafkaWriter := infrastructure.NewKafkaWriter(cfg.Kafka.Broker, cfg.Kafka.TopicPromise, true)

	ucHandler := usecase.HandlerConstructor{
		PromiseRepository: promiseRepo.CreateRepository(dbPool),
		EventProducer:     eventProducer.CreateProducer(kafkaWriter),
		ChunkSize:         cfg.App.ChunkSize,
		Metrics:           prometheusmetrics.NewServerMetrics(),
		Logger:            logger,
	}.New()

	httpServer := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: promhttp.Handler(),
	}

	listenerHTTP, err := net.Listen("tcp", cfg.HTTPServer.Addr)
	if err != nil {
		logger.Fatal(err.Error())
	}

	go func() {
		err := httpServer.Serve(listenerHTTP)
		if err != nil {
			logger.Fatal(err.Error())
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	logger.Info("starting http server", zap.String("addr", cfg.HTTPServer.Addr))

	grpcServer := infrastructure.InitGRPCServer(ucHandler, logger)

	grpc_prometheus.Register(grpcServer)

	listener, err := net.Listen("tcp", cfg.GRPCServer.Addr)
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("starting grpc server", zap.String("addr", cfg.GRPCServer.Addr))

	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatal(err.Error())
	}
}
