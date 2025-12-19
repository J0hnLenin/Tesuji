package gameservice

import (
	"context"

	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
)

func (s *GameService) GetGamesSummaryList(ctx context.Context, page uint64, limit uint64) ([]*proto_models.GameSummary, error) {
	return s.gameStorage.GetGamesSummaryList(ctx, page, limit)
}