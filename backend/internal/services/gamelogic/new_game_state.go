package gamelogic

import "github.com/J0hnLenin/Tesuji/internal/models"

func newGameState(m *models.Move, p models.Position) *models.GameState {
	return &models.GameState{
		LastMove: m,
		Position: p,
	}
}