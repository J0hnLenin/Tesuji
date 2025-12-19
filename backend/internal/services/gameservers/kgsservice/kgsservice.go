package kgsservice

import (
	"context"
	"net/http"
	"net/http/cookiejar"

	"github.com/J0hnLenin/Tesuji/config"
	"github.com/J0hnLenin/Tesuji/internal/models"
)

type GameService interface {
	UpsertGamesData (ctx context.Context, games []*models.Game) error
	InitGameData (ctx context.Context, game *models.Game) error
	MakeAMove (ctx context.Context, game *models.Game, move *models.Move) error
}

type KGSService struct {
	gameService GameService
	client *http.Client
	timeout uint16
	url    string
	user string
	pwd string
	srvID uint8
}

func NewKGSService(cfg *config.KGSConfig, gs GameService) *KGSService {
	jar, _ := cookiejar.New(nil)

	if cfg.SrvID < 10 || cfg.SrvID > 99 {
		panic("kgs.SrvID is not valid! Must be in [10, 99]")
	}

	return &KGSService{
		gameService: gs,
		client: &http.Client{
			Jar: jar,
		},
		timeout: cfg.Timeout,
		url: cfg.URL,
		user: cfg.User,
		pwd: cfg.Password,
		srvID: cfg.SrvID,
	}
}
