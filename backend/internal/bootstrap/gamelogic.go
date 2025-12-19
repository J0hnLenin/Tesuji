package bootstrap

import (
	"github.com/J0hnLenin/Tesuji/internal/services/gamelogic"
)

func InitGameLogic() *gamelogic.GameLogic {
	return gamelogic.NewGameLogic()
}