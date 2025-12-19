package bootstrap

import (
	"context"

	"github.com/J0hnLenin/Tesuji/config"
	"github.com/J0hnLenin/Tesuji/internal/producer/gameprod"
	"github.com/J0hnLenin/Tesuji/internal/services/gamelogic"
	"github.com/J0hnLenin/Tesuji/internal/services/gameservice"
	"github.com/J0hnLenin/Tesuji/internal/services/protoconvert"
	"github.com/J0hnLenin/Tesuji/internal/storage/pgstorage"
	"github.com/J0hnLenin/Tesuji/internal/storage/redisstorage"
)

func InitGameService(cfg *config.Config, st *pgstorage.PGstorage, rd *redisstorage.RedisStorage, gp *gameprod.GameProducer, gl *gamelogic.GameLogic, pc *protoconvert.ProtoConvert) *gameservice.GameService {
	return gameservice.NewGameService(context.Background(), st, rd, gp, gl, pc)
}