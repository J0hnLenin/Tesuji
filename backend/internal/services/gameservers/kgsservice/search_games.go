package kgsservice

import (
	"context"
	"log/slog"
	"time"
)

func (kgs *KGSService) SearchGames(ctx context.Context) {
	for {
		err := kgs.authenticate()
		if err != nil {
			slog.Error("KGSService.SearchGames authenticate error", "error", err.Error())
			continue
		}

		games, err := kgs.getGames(ctx)
		if err != nil {
			slog.Error("KGSService.SearchGames getGames error", "error", err.Error())
			continue
		}

		if len(games) > 0 {
			err = kgs.gameService.UpsertGamesData(ctx, games)
			if err != nil {
				slog.Error("KGSService.SearchGames saveGames error", "error", err.Error())
				continue
			}
		}
		time.Sleep(time.Duration(kgs.timeout) * time.Second)
	}
}