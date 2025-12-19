package gameservice

import (
	"context"

	"github.com/J0hnLenin/Tesuji/internal/models"
)

func (s *GameService) InitGameData(ctx context.Context, game *models.Game) error {
	return s.gameLogic.InitGameData(ctx, game)
}
