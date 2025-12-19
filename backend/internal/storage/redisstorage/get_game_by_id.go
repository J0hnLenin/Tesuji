package redisstorage

import (
	"context"
	"fmt"

	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
)

func (r *RedisStorage) GetGameByID(ctx context.Context, id int64) (*proto_models.Game, error) {

	key := getGameKey(id)

	gameData, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("game %d not found in Redis", id)
		}
		return nil, fmt.Errorf("failed to get game %d from Redis: %w", id, err)
	}

	game := &proto_models.Game{}
	if err := proto.Unmarshal(gameData, game); err != nil {
		return nil, fmt.Errorf("failed to unmarshal game %d: %w", id, err)
	}

	return game, nil
}