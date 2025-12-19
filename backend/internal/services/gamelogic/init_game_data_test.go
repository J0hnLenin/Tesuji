package gamelogic

import (
	"github.com/J0hnLenin/Tesuji/internal/models"
	"github.com/stretchr/testify/assert"
)

func (g *GameLogicSuite) TestInitGameDataEmptyError() {
	game := &models.Game{}

	err := g.gameLogic.InitGameData(g.ctx, game)

	assert.Error(g.T(), err)
}

func (g *GameLogicSuite) TestInitGameDataSuccess() {
	var wantGameSize uint8 = 19

	game := &models.Game{}
	game.BoardSize = wantGameSize

	err := g.gameLogic.InitGameData(g.ctx, game)
	gameSizeH := uint8(len(game.GameData[0].Position))
	gameSizeW := uint8(len(game.GameData[0].Position))

	assert.Nil(g.T(), err)
	assert.Equal(g.T(), 1, len(game.GameData))
	assert.Equal(g.T(), wantGameSize, gameSizeH)
	assert.Equal(g.T(), wantGameSize, gameSizeW)
}