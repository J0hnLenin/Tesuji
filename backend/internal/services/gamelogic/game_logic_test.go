package gamelogic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GameLogicSuite struct {
	suite.Suite
	ctx context.Context
	gameLogic *GameLogic
}

func (g *GameLogicSuite) SetupTest() {
	g.ctx = context.Background()
	g.gameLogic = NewGameLogic()
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(GameLogicSuite))
}