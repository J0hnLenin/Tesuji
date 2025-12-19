package gamelogic

import (
	"errors"

	"github.com/J0hnLenin/Tesuji/internal/models"
)

func addMoveOnPosition(pos models.Position, move *models.Move) error {
	if move.IsPass {
		return nil
	}
	if pos[move.Point.X][move.Point.Y] != models.EmptyPoint {
		return errors.New("can't put stone on not empty point")
	}
	pos[move.Point.X][move.Point.Y] = newStone(move.Color)
	neibors := getNeibors(pos, *move.Point)
	updatePointsStates(pos, neibors)
	return nil
}