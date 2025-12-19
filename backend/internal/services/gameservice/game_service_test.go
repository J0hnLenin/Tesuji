package gameservice

import (
	"context"
	"testing"

	"github.com/J0hnLenin/Tesuji/internal/services/gamelogic"
	mocks "github.com/J0hnLenin/Tesuji/internal/services/gameservice/mocks"
	"github.com/J0hnLenin/Tesuji/internal/services/protoconvert"
	"github.com/stretchr/testify/suite"
)

type GameServiceSuite struct {
	suite.Suite
	ctx            context.Context
	gameStorage *mocks.GameStorage
	gameCache   *mocks.GameCache
	gameProducer *mocks.GameProducer
	gameService *GameService
}

func (g *GameServiceSuite) SetupTest() {
	g.ctx = context.Background()
	g.gameStorage = mocks.NewGameStorage(g.T())
	g.gameCache = mocks.NewGameCache(g.T())
	g.gameProducer = mocks.NewGameProducer(g.T())
	gameLogic := gamelogic.NewGameLogic()
	protoConvert := protoconvert.NewProtoConvert()
	
	g.gameService = NewGameService(g.ctx, g.gameStorage, g.gameCache, g.gameProducer, gameLogic, protoConvert)
}
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(GameServiceSuite))
}
