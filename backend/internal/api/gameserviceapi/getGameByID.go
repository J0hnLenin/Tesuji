package gameserviceapi

import (
	"context"
	"log"

	"github.com/J0hnLenin/Tesuji/internal/pb/games_api"
)

func (g *GameServiceAPI) GetGameByID(ctx context.Context, req *games_api.GetGameByIDRequest) (*games_api.GetGameByIDResponse, error) {
	log.Printf("Received request with ID: %v", req.Id)

	responce, err := g.gameService.GetGameByID(ctx, req.Id)
	if err != nil {
		return &games_api.GetGameByIDResponse{}, err
	}
	return &games_api.GetGameByIDResponse{Game: responce}, nil
}