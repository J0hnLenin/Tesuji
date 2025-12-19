package pgstorage

import (
	"context"
	"fmt"

	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func (pg *PGstorage) GetGameByID(ctx context.Context, id int64) (*proto_models.Game, error) {
	query := squirrel.Select(gameProtoColumnName).
		From(tableWithBacket(pg.bucketByGameID(id))).
		Where(squirrel.Eq{idColumnName: id}).
		PlaceholderFormat(squirrel.Dollar)

	queryText, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "generate query error")
	}

	fmt.Println(queryText)

	var gameProto []byte
	
	res := pg.db.QueryRow(ctx, queryText, args...)

	if err = res.Scan(&gameProto); err != nil {
		return nil, errors.Wrap(err, "query row error")
	}

	game := &proto_models.Game{}
	if err := proto.Unmarshal(gameProto, game); err != nil {
		game.Summary.Id = id
		return nil, errors.Wrap(err, "unmarshal game proto error")
	}

	return game, nil
}