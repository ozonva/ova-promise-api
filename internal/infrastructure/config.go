package infrastructure

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	App struct {
		ChunkSize int `env:"APP_FLUSHER_CHUNK_SIZE" envDefault:"100"`
	}

	GRPCServer struct {
		Addr string `env:"GRPC_ADDR" envDefault:"127.0.0.1:9001"`
	}

	HTTPServer struct {
		Addr string `env:"HTTP_ADDR" envDefault:"127.0.0.1:9002"`
	}

	Database struct {
		ConnectionString string
		Host             string `env:"DB_HOST,required"`
		Port             int    `env:"DB_PORT,required"`
		User             string `env:"DB_USER,required"`
		Password         string `env:"DB_PASSWORD,required"`
		Name             string `env:"DB_NAME,required"`
		MaxConn          int32  `env:"DB_MAX_CONN" envDefault:"10"`
	}

	Kafka struct {
		Broker       string `env:"KAFKA_BROKER,required"`
		Group        string `env:"KAFKA_GROUP_ID,required"`
		TopicPromise string `env:"KAFKA_TOPIC_PROMISE,required"`
	}
}

func NewConfigFromEnv() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	cfg.Database.ConnectionString = fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	return cfg, nil
}
