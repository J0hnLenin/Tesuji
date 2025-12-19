package gameservice

import (
	"errors"

	"github.com/J0hnLenin/Tesuji/internal/models"
	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (g *GameServiceSuite) TestUpsertGamesDataSuccess() {
	
	game1 := &models.Game{}
	game1.BlackPlayer = &models.Player{}
	game1.WhitePlayer = &models.Player{}

	game2 := &models.Game{}
	game2.BlackPlayer = &models.Player{}
	game2.WhitePlayer = &models.Player{}
	game2.IsFinished = true

	games := []*models.Game{game1, game2}

	g.gameStorage.On("UpsertGamesData", g.ctx, mock.Anything).
		Return(nil).
		Once()

	g.gameCache.On("UpsertGamesData", g.ctx, mock.Anything).
		Return(nil).
		Once()

	g.gameProducer.On("ProduceGames", g.ctx, mock.Anything).
		Run(func(args mock.Arguments) {
			finishedGames := args.Get(1).([]*proto_models.Game)
			assert.Equal(g.T(), 1, len(finishedGames))
		}).
		Return(nil).
		Once()
	
	err := g.gameService.UpsertGamesData(g.ctx, games)

	assert.Nil(g.T(), err)
}

func (g *GameServiceSuite) TestUpsertGamesDataNothingToProduce() {
	
	game1 := &models.Game{}
	game1.BlackPlayer = &models.Player{}
	game1.WhitePlayer = &models.Player{}

	game2 := &models.Game{}
	game2.BlackPlayer = &models.Player{}
	game2.WhitePlayer = &models.Player{}

	games := []*models.Game{game1, game2}

	g.gameStorage.On("UpsertGamesData", g.ctx, mock.Anything).
		Return(nil).
		Once()

	g.gameCache.On("UpsertGamesData", g.ctx, mock.Anything).
		Return(nil).
		Once()

	err := g.gameService.UpsertGamesData(g.ctx, games)

	assert.Nil(g.T(), err)
}

func (g *GameServiceSuite) TestUpsertGamesDataProduceError() {
	
	wantErrorString := "Network Error"
	wantErr := errors.New(wantErrorString)

	game1 := &models.Game{}
	game1.BlackPlayer = &models.Player{}
	game1.WhitePlayer = &models.Player{}

	game2 := &models.Game{}
	game2.BlackPlayer = &models.Player{}
	game2.WhitePlayer = &models.Player{}
	game2.IsFinished = true

	games := []*models.Game{game1, game2}

	g.gameStorage.On("UpsertGamesData", g.ctx, mock.Anything).
		Return(nil).
		Once()

	g.gameCache.On("UpsertGamesData", g.ctx, mock.Anything).
		Return(nil).
		Once()

	g.gameProducer.On("ProduceGames", g.ctx, mock.Anything).
		Return(wantErr).
		Once()
	
	err := g.gameService.UpsertGamesData(g.ctx, games)

	assert.EqualError(g.T(), err, wantErrorString)
}

func (g *GameServiceSuite) TestUpsertGamesDataCacheError() {
	
	wantErrorString := "Network Error"
	wantErr := errors.New(wantErrorString)

	game1 := &models.Game{}
	game1.BlackPlayer = &models.Player{}
	game1.WhitePlayer = &models.Player{}

	game2 := &models.Game{}
	game2.BlackPlayer = &models.Player{}
	game2.WhitePlayer = &models.Player{}
	game2.IsFinished = true

	games := []*models.Game{game1, game2}

	g.gameStorage.On("UpsertGamesData", g.ctx, mock.Anything).
		Return(nil).
		Once()

	g.gameCache.On("UpsertGamesData", g.ctx, mock.Anything).
		Return(wantErr).
		Once()
	
	err := g.gameService.UpsertGamesData(g.ctx, games)

	assert.EqualError(g.T(), err, wantErrorString)
}

func (g *GameServiceSuite) TestUpsertGamesDataStorageError() {
	
	wantErrorString := "Network Error"
	wantErr := errors.New(wantErrorString)

	game1 := &models.Game{}
	game1.BlackPlayer = &models.Player{}
	game1.WhitePlayer = &models.Player{}

	game2 := &models.Game{}
	game2.BlackPlayer = &models.Player{}
	game2.WhitePlayer = &models.Player{}
	game2.IsFinished = true

	games := []*models.Game{game1, game2}

	g.gameStorage.On("UpsertGamesData", g.ctx, mock.Anything).
		Return(wantErr).
		Once()
	
	err := g.gameService.UpsertGamesData(g.ctx, games)

	assert.EqualError(g.T(), err, wantErrorString)
}