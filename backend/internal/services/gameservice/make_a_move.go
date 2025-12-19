package gameservice

import (
	"context"

	"github.com/J0hnLenin/Tesuji/internal/models"
)

func (s *GameService) MakeAMove(ctx context.Context, game *models.Game, move *models.Move) error {
	return s.gameLogic.MakeAMove(ctx, game, move)
}