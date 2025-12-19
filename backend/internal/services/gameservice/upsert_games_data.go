package gameservice

import (
	"context"
	"fmt"

	"github.com/J0hnLenin/Tesuji/internal/models"
)

func (s *GameService) UpsertGamesData(ctx context.Context, games []*models.Game) error {
	fmt.Println(len(games))
	for _, g := range games {
		fmt.Println(g.Title, " ", g.BlackPlayer.Name, " ", g.WhitePlayer.Name)
	}
	gamesProto := s.protoConvert.MapGamesToProto(games)

	err := s.gameStorage.UpsertGamesData(ctx, gamesProto)
	if err != nil {
		return err
	}

	err = s.gameCache.UpsertGamesData(ctx, gamesProto)
	if err != nil {
		return err
	}

	finishedGames := getFinishedGames(gamesProto)

	if len(finishedGames) == 0 {
		return nil
	}

	err = s.gameProducer.ProduceGames(ctx, finishedGames)
	if err != nil {
		return err
	}

	
	return nil
}
