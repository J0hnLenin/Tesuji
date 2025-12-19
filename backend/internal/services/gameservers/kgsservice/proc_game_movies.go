package kgsservice

import (
	"context"

	"github.com/J0hnLenin/Tesuji/internal/models"
)

func (kgs *KGSService) procGameMovies(ctx context.Context, game *models.Game, moves []*models.Move) error {

	err := kgs.gameService.InitGameData(ctx, game)
	if err != nil {
		return err
	}
	for _, move := range moves {
		err = kgs.gameService.MakeAMove(ctx, game, move)
		if err != nil {
			return err
		}
	}

	return nil
}