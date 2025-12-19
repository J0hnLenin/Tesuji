package kgsservice

import (
	"context"
	"log/slog"

	"github.com/J0hnLenin/Tesuji/internal/models"
)

func (kgs *KGSService) getGames(ctx context.Context) ([]*models.Game, error) {
	data, err := kgs.sendGetRequest()
	if err != nil {
		return nil, err
	}

	games, err := parseGames(data)
	if err != nil {
		return nil, err
	}
	var gamesInfo []*models.Game

	for _, game := range games {
		gameInfo, err := kgs.getGameInfo(ctx, game.ChannelID)
		if err != nil {
			slog.Error("KGSService.getGames getGameInfo error", "error", err.Error())
			continue
		}
		gamesInfo = append(gamesInfo, gameInfo)
	}

	return gamesInfo, err
}