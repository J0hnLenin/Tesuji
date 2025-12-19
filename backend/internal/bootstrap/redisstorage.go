package bootstrap

import (
	"fmt"
	"log"
	"strconv"

	"github.com/J0hnLenin/Tesuji/config"
	"github.com/J0hnLenin/Tesuji/internal/storage/redisstorage"
)

func InitRedisStorage(cfg *config.Config) *redisstorage.RedisStorage {

	db, err := strconv.Atoi(cfg.Redis.DBStr)
	if err != nil {
		db = 0
	}

	addr := fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)

	storage, err := redisstorage.NewRedisStorage(addr, cfg.Redis.Password, db)
	if err != nil {
		log.Panicf("ошибка инициализации БД, %v", err)
		panic(err)
	}
	return storage
}
