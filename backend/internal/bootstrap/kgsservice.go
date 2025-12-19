package bootstrap

import (
	"github.com/J0hnLenin/Tesuji/config"
	"github.com/J0hnLenin/Tesuji/internal/services/gameservers/kgsservice"
	"github.com/J0hnLenin/Tesuji/internal/services/gameservice"
)

func InitKGSService(cfg *config.Config, gs *gameservice.GameService) *kgsservice.KGSService {
	if cfg.GameServers.KGS.User == "" || cfg.GameServers.KGS.Password == "" || cfg.GameServers.KGS.URL == "" {
		panic("Empty KGS config")
	}
	return kgsservice.NewKGSService(&cfg.GameServers.KGS, gs)
}