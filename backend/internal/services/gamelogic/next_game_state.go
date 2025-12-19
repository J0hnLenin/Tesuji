package gamelogic

import "github.com/J0hnLenin/Tesuji/internal/models"

func nextGameState(gameState *models.GameState, newMove *models.Move) (*models.GameState, error) {
	newPosition := make(models.Position, len(gameState.Position))
    for i := range gameState.Position {
        newPosition[i] = make([]models.PointState, len(gameState.Position[i]))
        copy(newPosition[i], gameState.Position[i])
    }
	
	err := addMoveOnPosition(newPosition, newMove)
	if err != nil {
		return nil, err
	}
	return newGameState(newMove, newPosition), nil
}