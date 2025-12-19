package gamelogic

import (
	"context"
	"errors"

	"github.com/J0hnLenin/Tesuji/internal/models"
)

func (g *GameLogic) MakeAMove(ctx context.Context, game *models.Game, move *models.Move) error {
	if len(game.GameData) < 1 {
		return errors.New("cant make a move while gamedata not init")
	}
	currentGameState := &game.GameData[len(game.GameData)-1]
	gs, err := nextGameState(currentGameState, move)
	if err != nil {
		return err
	}
	game.GameData = append(game.GameData, *gs)
	return nil
}