package gameservice

import (
	"errors"

	"github.com/stretchr/testify/assert"
)

func (g *GameServiceSuite) TestGetGamesSummaryListSuccess() {
	
	var page uint64 = 1
	var limit uint64 = 10

	g.gameStorage.On("GetGamesSummaryList", g.ctx, page, limit).
		Return(nil, nil).
		Once()
	
	_, err := g.gameService.GetGamesSummaryList(g.ctx, page, limit)

	assert.Nil(g.T(), err)
}

func (g *GameServiceSuite) TestGetGamesSummaryListError() {
	
	wantErrorString := "Network Error"
	wantErr := errors.New(wantErrorString)

	var page uint64 = 1
	var limit uint64 = 10

	g.gameStorage.On("GetGamesSummaryList", g.ctx, page, limit).
		Return(nil, wantErr).
		Once()
	
	_, err := g.gameService.GetGamesSummaryList(g.ctx, page, limit)

	assert.EqualError(g.T(), err, wantErrorString)
}