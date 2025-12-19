package protoconvert

import (
	"time"

	"github.com/J0hnLenin/Tesuji/internal/models"
	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
	"github.com/samber/lo"
)

func (pc *ProtoConvert) MapGamesToProto(games []*models.Game) ([]*proto_models.Game) {
	return lo.Map(games, func(g *models.Game, _ int) *proto_models.Game {
		var gameData []*proto_models.GameState
		if g.GameData != nil {
			gameData = make([]*proto_models.GameState, len(g.GameData))
			for i, state := range g.GameData {
				gameData[i] = &proto_models.GameState{
					LastMove: mapMoveToProto(state.LastMove),
					Position: mapPositionToProto(state.Position),
				}
			}
		}

		return &proto_models.Game{
			Summary: &proto_models.GameSummary{
				Id:          int64(g.ID),
				BoardSize:   int32(g.BoardSize),
				Title:       g.Title,
				Komi:        g.Komi,
				Date:        g.Date.Format(time.RFC3339),
				IsFinished:  g.IsFinished,
				Rules:       proto_models.Rules(g.Rules),
				BlackPlayer: mapPlayerToProto(g.BlackPlayer),
				WhitePlayer: mapPlayerToProto(g.WhitePlayer),
			},
			GameData: gameData,
		}
	})
}