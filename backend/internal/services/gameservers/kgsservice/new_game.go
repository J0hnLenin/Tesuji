package kgsservice

import (
	"github.com/J0hnLenin/Tesuji/internal/models"
)

func newGame() *models.Game {
	black := &models.Player{}
	white := &models.Player{}
	game := &models.Game{
		GameSummary: models.GameSummary{
			BlackPlayer: black,
			WhitePlayer: white,
		},
	}
	return game
}