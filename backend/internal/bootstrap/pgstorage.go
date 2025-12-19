package bootstrap

import (
	"fmt"
	"log"

	"github.com/J0hnLenin/Tesuji/config"
	"github.com/J0hnLenin/Tesuji/internal/storage/pgstorage"
)

func InitPGStorage(cfg *config.Config) *pgstorage.PGstorage {

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)
	storage, err := pgstorage.NewPGStorge(connectionString, cfg.Database.BucketQuantity)
	if err != nil {
		log.Panicf("ошибка инициализации БД, %v", err)
		panic(err)
	}
	return storage
}
