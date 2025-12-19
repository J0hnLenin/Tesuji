package gameprod

import (
	"time"

	"github.com/segmentio/kafka-go"
)

type GameProducer struct {
	writer      *kafka.Writer
	kafkaBroker []string
	topicName   string
}

func NewGameProducer(kafkaBroker []string, topicName string) *GameProducer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(kafkaBroker...),
		Topic:        topicName,
		Balancer:     &kafka.Hash{},
		BatchTimeout: 10 * time.Millisecond,
		BatchSize:    100,
		RequiredAcks: kafka.RequireAll,
		Async:        false,
		Compression:  kafka.Snappy,
	}

	return &GameProducer{
		writer:      writer,
		kafkaBroker: kafkaBroker,
		topicName:   topicName,
	}
}
