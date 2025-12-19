package gamelogic

import "github.com/J0hnLenin/Tesuji/internal/models"

func updateState(pos models.Position, p models.Point) {
	needMarkDead, points := isDead(pos, p)
	if needMarkDead {
		markDead(pos, points)
	}
}
