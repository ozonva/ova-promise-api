package kafkaproducer

import (
	"github.com/segmentio/kafka-go"

	"github.com/ozonva/ova-promise-api/internal/usecase"
)

type w struct {
	writer *kafka.Writer
}

func CreateProducer(writer *kafka.Writer) usecase.EventProducer {
	return w{
		writer: writer,
	}
}
