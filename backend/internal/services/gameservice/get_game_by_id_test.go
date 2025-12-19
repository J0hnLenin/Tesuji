package gameservice

import (
	"errors"

	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
	"github.com/stretchr/testify/assert"
)

func (g *GameServiceSuite) TestGetGameByIDSuccessFromCache() {
	
	var id int64 = 1
	wantGame := &proto_models.Game{}
	wantGame.Summary = &proto_models.GameSummary{}
	wantGame.Summary.Id = id
	
	g.gameCache.On("GetGameByID", g.ctx, id).
		Return(wantGame, nil).
		Once()
	
	game, err := g.gameService.GetGameByID(g.ctx, id)

	assert.Nil(g.T(), err)
	assert.Equal(g.T(), wantGame, game)
}

func (g *GameServiceSuite) TestGetGameByIDSuccessFromStorage() {
	
	var id int64 = 1
	wantGame := &proto_models.Game{}
	wantGame.Summary = &proto_models.GameSummary{}
	wantGame.Summary.Id = id
	
	g.gameCache.On("GetGameByID", g.ctx, id).
		Return(nil, errors.New("Not found")).
		Once()

	g.gameStorage.On("GetGameByID", g.ctx, id).
		Return(wantGame, nil).
		Once()
	
	game, err := g.gameService.GetGameByID(g.ctx, id)

	assert.Nil(g.T(), err)
	assert.Equal(g.T(), wantGame, game)
}

func (g *GameServiceSuite) TestGetGameByIDSuccessError() {
	
	wantErrorString := "Network Error"
	wantErr := errors.New(wantErrorString)

	var id int64 = 1
	
	g.gameCache.On("GetGameByID", g.ctx, id).
		Return(nil, errors.New("Not found")).
		Once()

	g.gameStorage.On("GetGameByID", g.ctx, id).
		Return(nil, wantErr).
		Once()
	
	_, err := g.gameService.GetGameByID(g.ctx, id)

	assert.EqualError(g.T(), err, wantErrorString)
}