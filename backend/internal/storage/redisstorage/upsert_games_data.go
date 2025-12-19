package redisstorage

import (
	"context"
	"fmt"
	"time"

	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
	"google.golang.org/protobuf/proto"
)

func (r *RedisStorage) UpsertGamesData(ctx context.Context, games []*proto_models.Game) error {
	if len(games) == 0 {
		return nil
	}

	pipe := r.client.Pipeline()
	defer pipe.Discard()

	var errs []error

	for _, game := range games {
		if game == nil || game.Summary == nil {
			continue
		}

		gameData, err := proto.Marshal(game)
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to marshal game %d: %w", game.Summary.Id, err))
			continue
		}

		key := getGameKey(game.Summary.Id)

		pipe.Set(ctx, key, gameData, 0)

		pipe.HSet(ctx, "game:meta:"+string(game.Summary.Id),
			"updated_at", time.Now().Unix(),
			"is_finished", game.Summary.IsFinished,
		)
	}

	if len(errs) > 0 {
		return fmt.Errorf("serialization errors: %v", errs)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to execute pipeline: %w", err)
	}

	return nil
}