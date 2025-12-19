package gamelogic

import (
	"context"

	"github.com/J0hnLenin/Tesuji/internal/models"
)

func (g *GameLogic) InitGameData(ctx context.Context, game *models.Game) error {
	game.GameData = make([]models.GameState, 0)
	initPos, err := newPosition(game.BoardSize)
	if err != nil {
		return err
	}
	initGameState := newGameState(nil, initPos)
	game.GameData = append(game.GameData, *initGameState)
	return nil
}
