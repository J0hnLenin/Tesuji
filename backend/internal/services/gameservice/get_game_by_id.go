package gameservice

import (
	"context"
	"fmt"

	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
)

func (s *GameService) GetGameByID(ctx context.Context, id int64) (*proto_models.Game, error) {
	g, err := s.gameCache.GetGameByID(ctx, id)
	if err != nil {
		fmt.Print(err)
	}
	if g != nil {
		return g, nil
	}


	return s.gameStorage.GetGameByID(ctx, id)
}