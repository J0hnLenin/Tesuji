package gamelogic

import "github.com/J0hnLenin/Tesuji/internal/models"

func newStone(c models.Color) models.PointState {
	if c == models.Black {
		return models.BlackStone
	}
	return models.WhiteStone
}