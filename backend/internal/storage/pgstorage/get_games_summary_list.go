package pgstorage

import (
	"context"
	"fmt"
	"sort"
	"sync"

	proto_models "github.com/J0hnLenin/Tesuji/internal/pb/models"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
)

func (pg *PGstorage) GetGamesSummaryList(ctx context.Context, page uint64, limit uint64) ([]*proto_models.GameSummary, error) {
    if page < 1 {
        page = 1
    }
    offset := (page - 1) * limit
    
    g, ctx := errgroup.WithContext(ctx)
    var mu sync.Mutex
    allGames := make([]*proto_models.GameSummary, 0)
    
    for bucketID := bucketNum(0); bucketID < bucketNum(pg.bucketQuantity); bucketID++ {
        bucketID := bucketID
        g.Go(func() error {
            games, err := pg.bucketGamesSummary(ctx, limit, bucketID)
            if err != nil {
                return err
            }
            
            mu.Lock()
            allGames = append(allGames, games...)
            mu.Unlock()
            
            return nil
        })
    }
    
    if err := g.Wait(); err != nil {
        return nil, err
    }
    
    sort.Slice(allGames, func(i, j int) bool {
        return allGames[i].Date > allGames[j].Date
    })
    
    if len(allGames) == 0 || offset >= uint64(len(allGames)) {
        return []*proto_models.GameSummary{}, nil
    }
    
    end := offset + limit
    if end > uint64(len(allGames)) {
        end = uint64(len(allGames))
    }
    
    return allGames[offset:end], nil
}

func (pg *PGstorage) bucketGamesSummary(ctx context.Context, limit uint64, bucket bucketNum) ([]*proto_models.GameSummary, error) {
	query := squirrel.Select(idColumnName, gameProtoColumnName).
		From(tableWithBacket(bucket)).
		OrderBy(fmt.Sprintf("%s DESC", createdAtColumnName)).
		Limit(limit).
		PlaceholderFormat(squirrel.Dollar)

	queryText, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "generate query error")
	}

	fmt.Println(queryText)

	rows, err := pg.db.Query(ctx, queryText, args...)
	if err != nil {
		return nil, errors.Wrap(err, "query error")
	}
	defer rows.Close()

	var summaries []*proto_models.GameSummary
	for rows.Next() {
		var gameProto []byte
		var id int64

		if err := rows.Scan(&id, &gameProto); err != nil {
			return nil, errors.Wrap(err, "scan row error")
		}

		game := &proto_models.Game{}
		if err := proto.Unmarshal(gameProto, game); err != nil {
			return nil, errors.Wrap(err, "unmarshal game proto error")
		}

		if game.Summary != nil {
			game.Summary.Id = id
			summaries = append(summaries, game.Summary)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows iteration error")
	}

	return summaries, nil
}