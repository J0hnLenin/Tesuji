package main

import (
	"fmt"
	"os"

	"github.com/J0hnLenin/Tesuji/config"
	"github.com/J0hnLenin/Tesuji/internal/bootstrap"
)

func main() {

	cfg, err := config.LoadConfig(os.Getenv("configPath"))
	if err != nil {
		panic(fmt.Sprintf("ошибка парсинга конфига, %v", err))
	}

	gameStorage := bootstrap.InitPGStorage(cfg)
	redisStorage := bootstrap.InitRedisStorage(cfg)
	gameProducer := bootstrap.InitGameProducer(cfg)
	gameLogic := bootstrap.InitGameLogic()
	protoConvert := bootstrap.InitProtoCnvert()
	gameService := bootstrap.InitGameService(cfg, gameStorage, redisStorage, gameProducer, gameLogic, protoConvert)
	kgsService := bootstrap.InitKGSService(cfg, gameService)
	gamesApi := bootstrap.InitGameServiceAPI(gameService)

	bootstrap.AppRun(kgsService, *gamesApi)
}
