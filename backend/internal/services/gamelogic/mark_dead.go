package gamelogic

import "github.com/J0hnLenin/Tesuji/internal/models"

func markDead(pos models.Position, points map[models.Point]bool) {
	for p := range points {
		pos[p.X][p.Y] = models.EmptyPoint
	}
}