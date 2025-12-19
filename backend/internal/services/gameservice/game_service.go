package gameservice

import (
	"context"

	"github.com/J0hnLenin/Tesuji/internal/models"
	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
)

type gameLogic interface {
	InitGameData (ctx context.Context, game *models.Game) error
	MakeAMove (ctx context.Context, game *models.Game, move *models.Move) error
}

type protoConvert interface {
	MapGamesToProto (games []*models.Game) []*proto_models.Game
}

//go:generate mockery --name gameStorage
type GameStorage interface {
	GetGameByID (ctx context.Context, id int64) (*proto_models.Game, error)
	GetGamesSummaryList (ctx context.Context, page uint64, limit uint64) ([]*proto_models.GameSummary, error)
	UpsertGamesData (ctx context.Context, games []*proto_models.Game) error
}

//go:generate mockery --name gameCache
type GameCache interface {
	GetGameByID (ctx context.Context, id int64) (*proto_models.Game, error)
	UpsertGamesData (ctx context.Context, games []*proto_models.Game) error
}

//go:generate mockery --name gameProducer
type GameProducer interface {
	ProduceGames (ctx context.Context, games []*proto_models.Game) error
}

type GameService struct {
	gameStorage  GameStorage
	gameCache    GameCache
	gameProducer GameProducer
	gameLogic    gameLogic
	protoConvert protoConvert
}

func NewGameService(ctx context.Context, gs GameStorage, gc GameCache, gp GameProducer, gl gameLogic, pc protoConvert) *GameService {
	return &GameService{
		gameStorage: gs,
		gameCache: gc,
		gameProducer: gp,
		gameLogic: gl,
		protoConvert: pc,
	}
}