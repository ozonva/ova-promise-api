package infrastructure

import (
	"github.com/segmentio/kafka-go"
)

func NewKafkaWriter(addr, topic string, async bool) *kafka.Writer {
	return &kafka.Writer{
		Addr:  kafka.TCP(addr),
		Topic: topic,
		Async: async,
	}
}
