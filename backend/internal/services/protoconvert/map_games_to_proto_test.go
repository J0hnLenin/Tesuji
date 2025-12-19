package protoconvert

import (
	"time"

	"github.com/J0hnLenin/Tesuji/internal/models"
	"github.com/stretchr/testify/assert"
)

func (p *ProtoConvertSuite) TestMapGamesToProtoSuccess() {
	now := time.Now()
	
	testGame := &models.Game{
		GameSummary: models.GameSummary{
			ID:         123,
			BoardSize:  3,
			Title:      "Test Game",
			Komi:       6.5,
			Date:       now,
			IsFinished: true,
			Rules:      models.Japanese,
			BlackPlayer: &models.Player{
				Name: "Black Player",
				Rank: 1,
			},
			WhitePlayer: &models.Player{
				Name: "White Player",
				Rank: 2,
			},
		},
		GameData: []models.GameState{
			{
				LastMove: &models.Move{
					Point: &models.Point{
						X: 1,
						Y: 2,
					},
					Color:    models.Black,
					IsPass:   false,
				},
				Position: models.Position{
						{
							models.BlackStone,
							models.EmptyPoint,
							models.WhiteStone,
						},
						{
							models.BlackStone,
							models.BlackStone,
							models.WhiteStone,
						},
						{
							models.WhiteStone,
							models.WhiteStone,
							models.WhiteStone,
						},
					},
			},
			{
				LastMove: &models.Move{
					Point: 	  nil,
					Color:    models.White,
					IsPass:   true,
				},
				Position: models.Position{
						{
							models.BlackStone,
							models.EmptyPoint,
							models.WhiteStone,
						},
						{
							models.BlackStone,
							models.EmptyPoint,
							models.WhiteStone,
						},
						{
							models.WhiteStone,
							models.WhiteStone,
							models.WhiteStone,
						},
					},
			},
			{
				LastMove: nil,
				Position: models.Position{},
			},
			{
				LastMove: nil,
				Position: nil,
			},
		},
	}
	emptyGame := &models.Game{}
	

	games := []*models.Game{testGame, emptyGame}
	result := p.protoConvert.MapGamesToProto(games)

	assert.Equal(p.T(), 2, len(result))
	
	protoGame := result[0]
	
	assert.NotNil(p.T(), protoGame.Summary)
	assert.Equal(p.T(), int64(123), protoGame.Summary.Id)
	assert.Equal(p.T(), int32(3), protoGame.Summary.BoardSize)
	assert.Equal(p.T(), testGame.Title, protoGame.Summary.Title)
	assert.Equal(p.T(), float32(6.5), protoGame.Summary.Komi)
	assert.Equal(p.T(), now.Format(time.RFC3339), protoGame.Summary.Date)
	assert.Equal(p.T(), true, protoGame.Summary.IsFinished)
	assert.Equal(p.T(), string(models.Japanese), string(protoGame.Summary.Rules))
	
	assert.NotNil(p.T(), protoGame.Summary.BlackPlayer)
	assert.Equal(p.T(), testGame.BlackPlayer.Name, protoGame.Summary.BlackPlayer.Name)
	
	assert.NotNil(p.T(), protoGame.Summary.WhitePlayer)
	assert.Equal(p.T(), testGame.WhitePlayer.Name, protoGame.Summary.WhitePlayer.Name)
	
	assert.Equal(p.T(), len(testGame.GameData), len(protoGame.GameData))
	assert.NotNil(p.T(), protoGame.GameData[0])
	assert.NotNil(p.T(), protoGame.GameData[0].LastMove)
	assert.NotNil(p.T(), protoGame.GameData[0].Position)

	assert.NotEqual(p.T(), protoGame.GameData[0].Position, protoGame.GameData[1].Position)
	assert.NotEqual(p.T(), protoGame.GameData[1].Position, protoGame.GameData[2].Position)
}

func (p *ProtoConvertSuite) TestMapGamesToProtoEmptySlice() {
	games := []*models.Game{}
	result := p.protoConvert.MapGamesToProto(games)
	
	assert.Equal(p.T(), 0, len(result))
	assert.NotNil(p.T(), result)
}