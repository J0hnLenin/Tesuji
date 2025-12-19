package gameserviceapi

import (
	"context"
	"log"

	"github.com/J0hnLenin/Tesuji/internal/pb/games_api"
)

func (g *GameServiceAPI) GetGamesSummaryList(ctx context.Context, req *games_api.GetGamesSummaryListRequest) (*games_api.GetGamesSummaryListResponse, error) {
	log.Printf("Received request with page: %v", req.Page)

	responce, err := g.gameService.GetGamesSummaryList(ctx, uint64(req.Page), uint64(req.Limit))
	if err != nil {
		return &games_api.GetGamesSummaryListResponse{}, err
	}
	return &games_api.GetGamesSummaryListResponse{Games: responce}, nil
}