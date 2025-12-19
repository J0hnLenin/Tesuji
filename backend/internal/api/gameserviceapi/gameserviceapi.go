package gameserviceapi

import (
	"context"

	"github.com/J0hnLenin/Tesuji/internal/pb/games_api"
	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
)

type gameService interface {
	GetGameByID(ctx context.Context, ID int64) (*proto_models.Game, error)
	GetGamesSummaryList(ctx context.Context, page uint64, limit uint64) ([]*proto_models.GameSummary, error)
}

type GameServiceAPI struct {
	games_api.UnimplementedGamesServiceServer
	gameService gameService
}

func NewGameServiceAPI(s gameService) *GameServiceAPI {
	return &GameServiceAPI{
		gameService: s,
	}
}
