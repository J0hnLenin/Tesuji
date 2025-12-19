package kgsservice

import (
	"context"
	"fmt"
	"strconv"

	"github.com/J0hnLenin/Tesuji/internal/models"
)

func (kgs *KGSService) getGameInfo(ctx context.Context, channelID int) (*models.Game, error) {

	joinReq := map[string]interface{}{
		"type":      "JOIN_REQUEST",
		"channelId": channelID,
	}

	err := kgs.sendPostRequest(joinReq)
	if err != nil {
		return nil, err
	}

	res, err := kgs.sendGetRequest()
	if err != nil {
		return nil, err
	}

	game, moves, err := parseGameMoves(res)
	if err != nil {
		return nil, err
	}

	err = kgs.procGameMovies(ctx, game, moves)
	if err != nil {
		return nil, err
	}

	game.ID, err = strconv.ParseUint(fmt.Sprint(kgs.srvID) + fmt.Sprint(channelID), 10, 64)
	if err != nil {
		return nil, err
	} 
	
	return game, nil
}