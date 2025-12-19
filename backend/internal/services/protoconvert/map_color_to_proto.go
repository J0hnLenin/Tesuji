package protoconvert

import (
	"github.com/J0hnLenin/Tesuji/internal/models"
	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
)

func mapColorToProto(c *models.Color) *proto_models.Color {
	if *c == models.Black {
		return proto_models.Color_BLACK.Enum()
	}
	if *c == models.White {
		return proto_models.Color_WHITE.Enum()
	}
	return nil
}