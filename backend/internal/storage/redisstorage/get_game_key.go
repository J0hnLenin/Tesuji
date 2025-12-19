package redisstorage

import (
	"fmt"
)
func getGameKey(gameID int64) string {
	return fmt.Sprintf("g:%d", gameID)
}