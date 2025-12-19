package protoconvert

import (
	"github.com/J0hnLenin/Tesuji/internal/models"
	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
)

func mapMoveToProto(m *models.Move) *proto_models.Move {
	if m == nil {
		return nil
	}

	color := mapColorToProto(&(m.Color))

	move := &proto_models.Move{
		IsPass: m.IsPass,
		Color:  proto_models.Color(*color),
	}

	if m.Point != nil {
		move.Point = &proto_models.Point{
			X: int32(m.Point.X),
			Y: int32(m.Point.Y),
		}
	}

	return move
}