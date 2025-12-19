package bootstrap

import (
	server "github.com/J0hnLenin/Tesuji/internal/api/gameserviceapi"
	"github.com/J0hnLenin/Tesuji/internal/services/gameservice"
)

func InitGameServiceAPI(gameService *gameservice.GameService) *server.GameServiceAPI {
	return server.NewGameServiceAPI(gameService)
}