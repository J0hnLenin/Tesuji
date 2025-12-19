package pgstorage

import (
	"fmt"

	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
	"github.com/Masterminds/squirrel"
	"github.com/samber/lo"
	"google.golang.org/protobuf/proto"
)

func upsertQuery(games []*proto_models.Game, bucket bucketNum) (squirrel.Sqlizer, error) {
	var err error
	infos := lo.Map(games, func(g *proto_models.Game, _ int) *game {
		var gameBytes []byte
		gameBytes, err = proto.Marshal(g)
		return &game{
			ID:        g.Summary.Id,
			GameProto: gameBytes,
		}
	})

	if err != nil {
		return nil, err
	}

	tableWithBacket :=tableWithBacket(bucket)

	q := squirrel.Insert(tableWithBacket).Columns(idColumnName, gameProtoColumnName).
		PlaceholderFormat(squirrel.Dollar)

	for _, info := range infos {
		q = q.Values(info.ID, info.GameProto)
	}

	q = q.Suffix(
		fmt.Sprintf("ON CONFLICT (%s) DO UPDATE SET %s = EXCLUDED.%s, %s = CURRENT_TIMESTAMP",
			idColumnName,
			gameProtoColumnName,
			gameProtoColumnName,
			updatedAtColumnName,
		),
	)

	return q, nil
}