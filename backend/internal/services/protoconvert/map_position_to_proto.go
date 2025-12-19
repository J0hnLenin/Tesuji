package protoconvert

import (
	"github.com/J0hnLenin/Tesuji/internal/models"
	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
)

func mapPositionToProto(pos models.Position) *proto_models.Position {
	n := len(pos)
	if n == 0 {
		return nil
	}

	data := make([]int32, n*n)
	for i, row := range pos {
		for j, val := range row {
			data[n*i + j] = int32(val)
		}
	}

	return &proto_models.Position{
		Data: data,
	}
}