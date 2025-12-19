package gamelogic

import (
	"errors"

	"github.com/J0hnLenin/Tesuji/internal/models"
)

func newPosition(size uint8) (models.Position, error) {
	if size == 0 {
		return nil, errors.New("0 is invalid board size")
	}
	p := make(models.Position, size)
	for i := range p {
		p[i] = make([]models.PointState, size)
	}
	return p, nil
}