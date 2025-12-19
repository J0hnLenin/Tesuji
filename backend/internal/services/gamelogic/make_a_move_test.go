package gamelogic

import (
	"github.com/J0hnLenin/Tesuji/internal/models"
	"github.com/stretchr/testify/assert"
)

func (g *GameLogicSuite) TestMakeAMoveSuccess() {
	var wantPosX uint8 = 1
	var wantPosY uint8 = 2

	wantMoves := 2
	wantColor := models.Black
	wantPointState := models.BlackStone

	game := &models.Game{}
	game.BoardSize = 19
	g.gameLogic.InitGameData(g.ctx, game)

	move := &models.Move{
		Color: wantColor,
		IsPass: false,
		Point: &models.Point{
			X: wantPosX,
			Y: wantPosY,
		},
	}

	err := g.gameLogic.MakeAMove(g.ctx, game, move)
	
	assert.Nil(g.T(), err)
	assert.Equal(g.T(), wantMoves, len(game.GameData))
	assert.Equal(g.T(), wantPosX, game.GameData[1].LastMove.Point.X)
	assert.Equal(g.T(), wantPosY, game.GameData[1].LastMove.Point.Y)
	assert.Equal(g.T(), wantColor, game.GameData[1].LastMove.Color)
	assert.Equal(g.T(), wantPointState, game.GameData[1].Position[wantPosX][wantPosY])
	assert.Equal(g.T(), models.EmptyPoint, game.GameData[0].Position[wantPosX][wantPosY])
}

func (g *GameLogicSuite) TestMakeAMoveErrorDoubleMove() {
	var wantPosX uint8 = 1
	var wantPosY uint8 = 2

	wantMoves := 2
	wantColor := models.Black

	game := &models.Game{}
	game.BoardSize = 19
	g.gameLogic.InitGameData(g.ctx, game)

	move := &models.Move{
		Color: wantColor,
		IsPass: false,
		Point: &models.Point{
			X: wantPosX,
			Y: wantPosY,
		},
	}

	g.gameLogic.MakeAMove(g.ctx, game, move)
	err := g.gameLogic.MakeAMove(g.ctx, game, move)

	assert.NotNil(g.T(), err)
	assert.Equal(g.T(), wantMoves, len(game.GameData))
}

func (g *GameLogicSuite) TestMakeAMoveErrorGameNotInit() {
	game := &models.Game{}
	game.BoardSize = 19

	move := &models.Move{
		Color: models.Black,
		IsPass: false,
		Point: &models.Point{
			X: 0,
			Y: 0,
		},
	}

	err := g.gameLogic.MakeAMove(g.ctx, game, move)

	assert.NotNil(g.T(), err)
}

func (g *GameLogicSuite) TestMakeAMoveSuccessCapture() {
	// From:
	//   0 1 2
	// 0 . b w
	// 1 b w .
	// 2 . b .
	//
	// Want:
	//   0 1 2
	// 0 . b .
	// 1 b . b
	// 2 . b .

	game := &models.Game{}
	game.BoardSize = 3
	g.gameLogic.InitGameData(g.ctx, game)

	moves := []*models.Move{
		{
			Color: models.Black,
			IsPass: false,
			Point: newPoint(0, 1),
		},
		{
			Color: models.White,
			IsPass: false,
			Point: newPoint(0, 2),
		},
		{
			Color: models.Black,
			IsPass: false,
			Point: newPoint(1, 0),
		},
		{
			Color: models.White,
			IsPass: false,
			Point: newPoint(1, 1),
		},
		{
			Color: models.Black,
			IsPass: false,
			Point: newPoint(2, 1),
		},
		{
			Color: models.White,
			IsPass: true,
			Point: nil,
		},
		{
			Color: models.Black,
			IsPass: false,
			Point: newPoint(1, 2),
		},
	}

	for _, move := range moves {
		err := g.gameLogic.MakeAMove(g.ctx, game, move)
		assert.Nil(g.T(), err)
	}
	
	pos := game.GameData[len(game.GameData)-1].Position
	assert.Equal(g.T(), models.EmptyPoint, pos[0][0])
	assert.Equal(g.T(), models.BlackStone, pos[0][1])
	assert.Equal(g.T(), models.EmptyPoint, pos[0][2])

	assert.Equal(g.T(), models.BlackStone, pos[1][0])
	assert.Equal(g.T(), models.EmptyPoint, pos[1][1])
	assert.Equal(g.T(), models.BlackStone, pos[1][2])

	assert.Equal(g.T(), models.EmptyPoint, pos[2][0])
	assert.Equal(g.T(), models.BlackStone, pos[2][1])
	assert.Equal(g.T(), models.EmptyPoint, pos[2][2])
}

func (g *GameLogicSuite) TestMakeAMoveSuccessAnotherCapture() {
	// From:
	//   0 1
	// 0 b b
	// 1 b .
	//
	// Want:
	//   0 1
	// 0 . .
	// 1 . w

	game := &models.Game{}
	game.BoardSize = 2
	g.gameLogic.InitGameData(g.ctx, game)

	moves := []*models.Move{
		{
			Color: models.Black,
			IsPass: false,
			Point: newPoint(0, 0),
		},
		{
			Color: models.White,
			IsPass: true,
			Point: nil,
		},
		{
			Color: models.Black,
			IsPass: false,
			Point: newPoint(0, 1),
		},
		{
			Color: models.White,
			IsPass: true,
			Point: nil,
		},
		{
			Color: models.Black,
			IsPass: false,
			Point: newPoint(1, 0),
		},
		{
			Color: models.White,
			IsPass: false,
			Point: newPoint(1, 1),
		},
	}

	for _, move := range moves {
		err := g.gameLogic.MakeAMove(g.ctx, game, move)
		assert.Nil(g.T(), err)
	}
	
	pos := game.GameData[len(game.GameData)-1].Position
	assert.Equal(g.T(), models.EmptyPoint, pos[0][0])
	assert.Equal(g.T(), models.EmptyPoint, pos[0][1])

	assert.Equal(g.T(), models.EmptyPoint, pos[1][0])
	assert.Equal(g.T(), models.WhiteStone, pos[1][1])
}

func (g *GameLogicSuite) TestMakeAMoveSuccessNotCapture() {
	// From:
	//   0 1
	// 0 b b
	// 1 . .
	//
	// Want:
	//   0 1
	// 0 b b
	// 1 . w

	game := &models.Game{}
	game.BoardSize = 2
	g.gameLogic.InitGameData(g.ctx, game)

	moves := []*models.Move{
		{
			Color: models.Black,
			IsPass: false,
			Point: newPoint(0, 0),
		},
		{
			Color: models.White,
			IsPass: true,
			Point: nil,
		},
		{
			Color: models.Black,
			IsPass: false,
			Point: newPoint(0, 1),
		},
		{
			Color: models.White,
			IsPass: false,
			Point: newPoint(1, 1),
		},
	}

	for _, move := range moves {
		err := g.gameLogic.MakeAMove(g.ctx, game, move)
		assert.Nil(g.T(), err)
	}
	
	pos := game.GameData[len(game.GameData)-1].Position
	assert.Equal(g.T(), models.BlackStone, pos[0][0])
	assert.Equal(g.T(), models.BlackStone, pos[0][1])

	assert.Equal(g.T(), models.EmptyPoint, pos[1][0])
	assert.Equal(g.T(), models.WhiteStone, pos[1][1])
}
