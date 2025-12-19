package pgstorage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type PGstorage struct {
	db *pgxpool.Pool
	bucketQuantity uint16
}

func NewPGStorge(connString string, bucketQuantity uint16) (*PGstorage, error) {

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, errors.Wrap(err, "ошибка парсинга конфига")
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, errors.Wrap(err, "ошибка подключения")
	}

	if bucketQuantity == 0 {
		return nil, errors.Wrap(err, "количество шардов Postgres должно быть больше нуля")
	}
	storage := &PGstorage{
		db: db,
		bucketQuantity: bucketQuantity,
	}

	return storage, nil
}
