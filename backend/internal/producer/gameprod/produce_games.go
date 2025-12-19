package gameprod

import (
	"context"
	"fmt"
	"log/slog"

	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
)


func (p *GameProducer) ProduceGames(ctx context.Context, gamesProto []*proto_models.Game) error {
	if len(gamesProto) == 0 {
		return nil
	}

	var firstError error
	successCount := 0

	for _, game := range gamesProto {	
		err := p.ProduceGame(ctx, game)
		if err != nil {
			slog.Error("Failed to produce game in batch",
				"game_id", game.Summary.Id,
				"error", err)
			
			if firstError == nil {
				firstError = fmt.Errorf("failed to produce game %d: %w", game.Summary.Id, err)
			}
		} else {
			successCount++
		}
	}

	slog.Info("Batch production completed",
		"total", len(gamesProto),
		"success", successCount,
		"failed", len(gamesProto)-successCount)

	if firstError != nil && successCount == 0 {
		return firstError
	}
	
	if firstError != nil {
		return fmt.Errorf("partial failure: %w (success: %d/%d)", 
			firstError, successCount, len(gamesProto))
	}

	return nil
}