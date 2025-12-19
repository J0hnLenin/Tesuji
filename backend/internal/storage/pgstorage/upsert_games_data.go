package pgstorage

import (
	"context"
	"fmt"

	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
	"github.com/pkg/errors"
)

func (pg *PGstorage) UpsertGamesData(ctx context.Context, games []*proto_models.Game) error {
    gamesByBucket := pg.groupGamesByBucket(games)

    for bucket, bucketGames := range gamesByBucket {
        query, err := upsertQuery(bucketGames, bucket)
        if err != nil {
            return errors.Wrapf(err, "proto.Marshal error for bucket %d", bucket)
        }
        
        queryText, args, err := query.ToSql()
        if err != nil {
            return errors.Wrapf(err, "generate query error for bucket %d", bucket)
        }
        
        fmt.Println(queryText)
        _, err = pg.db.Exec(ctx, queryText, args...)
        if err != nil {
            return errors.Wrapf(err, "exec query for bucket %d", bucket)
        }
    }
    
    return nil
}