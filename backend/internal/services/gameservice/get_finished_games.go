package gameservice

import proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"

func getFinishedGames(games []*proto_models.Game) []*proto_models.Game {
	var finishedGames []*proto_models.Game

	for _, game := range games {
		if game != nil && game.Summary != nil && game.Summary.IsFinished {
			finishedGames = append(finishedGames, game)
		}
	}

	return finishedGames
}