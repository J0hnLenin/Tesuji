package pgstorage

import proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"

func (pg *PGstorage) groupGamesByBucket(games []*proto_models.Game) map[bucketNum][]*proto_models.Game {
	gamesByBucket := make(map[bucketNum][]*proto_models.Game)
	for _, game := range games {
		bucket := pg.bucketByGameID(game.Summary.Id)
		gamesByBucket[bucket] = append(gamesByBucket[bucket], game)
	}
	return gamesByBucket
}