package protoconvert

import (
	"github.com/J0hnLenin/Tesuji/internal/models"
	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
)

func mapPlayerToProto(p *models.Player) *proto_models.Player {
	if p == nil {
		return nil
	}
	return &proto_models.Player{
		Name: p.Name,
		Rank: int32(p.Rank),
	}
}