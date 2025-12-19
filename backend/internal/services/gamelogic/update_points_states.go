package gamelogic

import "github.com/J0hnLenin/Tesuji/internal/models"

func updatePointsStates(pos models.Position, points []models.Point) {
	for _, p := range points {
		updateState(pos, p)
	}
}
