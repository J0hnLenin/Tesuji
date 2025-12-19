package gameservice

import (
	"github.com/J0hnLenin/Tesuji/internal/models"
	"github.com/stretchr/testify/assert"
)

func (g *GameServiceSuite) TestMakeAMoveSuccess() {
	var wantPosX uint8 = 1
	var wantPosY uint8 = 2

	wantMoves := 2
	wantColor := models.Black
	wantPointState := models.BlackStone

	game := &models.Game{}
	game.BoardSize = 19
	g.gameService.InitGameData(g.ctx, game)

	move := &models.Move{
		Color: wantColor,
		IsPass: false,
		Point: &models.Point{
			X: wantPosX,
			Y: wantPosY,
		},
	}

	err := g.gameService.MakeAMove(g.ctx, game, move)
	
	assert.Nil(g.T(), err)
	assert.Equal(g.T(), wantMoves, len(game.GameData))
	assert.Equal(g.T(), wantPosX, game.GameData[1].LastMove.Point.X)
	assert.Equal(g.T(), wantPosY, game.GameData[1].LastMove.Point.Y)
	assert.Equal(g.T(), wantColor, game.GameData[1].LastMove.Color)
	assert.Equal(g.T(), wantPointState, game.GameData[1].Position[wantPosX][wantPosY])
}

func (g *GameServiceSuite) TestMakeAMoveErrorDoubleMove() {
	var wantPosX uint8 = 1
	var wantPosY uint8 = 2

	wantMoves := 2
	wantColor := models.Black

	game := &models.Game{}
	game.BoardSize = 19
	g.gameService.InitGameData(g.ctx, game)

	move := &models.Move{
		Color: wantColor,
		IsPass: false,
		Point: &models.Point{
			X: wantPosX,
			Y: wantPosY,
		},
	}

	g.gameService.MakeAMove(g.ctx, game, move)
	err := g.gameService.MakeAMove(g.ctx, game, move)

	assert.NotNil(g.T(), err)
	assert.Equal(g.T(), wantMoves, len(game.GameData))
}
