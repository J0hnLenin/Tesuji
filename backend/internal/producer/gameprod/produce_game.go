package gameprod

import (
	"context"
	"log/slog"
	"time"

	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)


func (p *GameProducer) ProduceGame(ctx context.Context, gameProto *proto_models.Game) error {

	value, err := proto.Marshal(gameProto)
	if err != nil {
		slog.Error("Failed to marshal game proto", "game_id", gameProto.Summary.Id, "error", err)
		return err
	}

	message := kafka.Message{
		Key:   []byte(string(rune(gameProto.Summary.Id))),
		Value: value,
		Time:  time.Now(),
		Headers: []kafka.Header{
			{Key: "event-type", Value: []byte("game-upsert-proto")},
			{Key: "content-type", Value: []byte("application/protobuf")},
		},
	}

	err = p.writer.WriteMessages(ctx, message)
	if err != nil {
		slog.Error("Failed to write game proto to Kafka",
			"topic", p.topicName,
			"game_id", gameProto.Summary.Id,
			"error", err)
		return err
	}

	slog.Debug("Successfully produced game proto to Kafka",
		"topic", p.topicName,
		"game_id", gameProto.Summary.Id)

	return nil
}