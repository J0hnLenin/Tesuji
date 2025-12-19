package gamelogic

import "github.com/J0hnLenin/Tesuji/internal/models"

func newPoint(x uint8, y uint8) *models.Point {
	return &models.Point{
		X: x,
		Y: y,
	}
}