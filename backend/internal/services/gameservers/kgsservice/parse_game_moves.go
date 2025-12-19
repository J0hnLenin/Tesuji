package kgsservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/J0hnLenin/Tesuji/internal/models"
)

func parseGameMoves(jsonData string) (*models.Game, []*models.Move, error) {
	var response response
	game := newGame()

	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		fmt.Println(jsonData)
		return nil, nil, err
	}

	moves := make([]*models.Move, 0)
	for _, message := range response.Messages {
		if message.Type == "PRIVATE_KEEP_OUT" {
			return nil, nil, errors.New("cant connect to private game")
		}
		if message.Type == "GAME_JOIN" && len(message.SgfEvents) > 0 {
			for _, event := range message.SgfEvents {
				if event.Type == "PROP_GROUP_ADDED" {
					for _, prop := range event.Props {
						switch prop.Name {
						case "GAMENAME":
							game.Title = prop.Text
						case "DATE":
							game.Date = time.Now()
						case "RULES":
							game.BoardSize = prop.Size
							game.Komi = prop.Komi
							if prop.Rules == "japanese" {
								game.Rules = models.Japanese
							}
						case "PLAYERNAME":
							switch prop.Color {
							case "black":
								game.BlackPlayer.Name = prop.Text
							case "white":
								game.WhitePlayer.Name = prop.Text
							}
						case "PLAYERRANK":
							switch prop.Color {
							case "black":
								game.BlackPlayer.Rank = uint8(prop.Int)
							case "white":
								game.WhitePlayer.Rank = uint8(prop.Int)
							}
						case "MOVE":
							move := models.Move{}

							switch prop.Color {
							case "black":
								move.Color = models.Black
							case "white":
								move.Color = models.White
							}

							if isPass(&prop) {
								move.IsPass = true
							} else {
								x, y, success := getCoordinates(&prop)
								if !success {
									fmt.Println("cant parse move: ", prop)
									continue
								}
								point := models.Point{
									X: uint8(x),
									Y: uint8(y),
								}
								move.Point = &point
								move.IsPass = false
							}

							moves = append(moves, &move)
						case "RESULT":
							game.IsFinished = true
						case "TIMELEFT":
							continue
						case "PLACE":
							continue
						case "COMMENT":
							continue
						default: 
							fmt.Println("Unknown prop: ", prop)
						}
					}
				}
			}
		}
	}

	return game, moves, nil
}

func isPass(p *property) bool {
    if p.Loc == nil {
        return false
    }
    if str, ok := p.Loc.(string); ok {
        return str == "PASS"
    }
    return false
}

func getCoordinates(p *property) (int, int, bool) {
    if p.Loc == nil {
        return 0, 0, false
    }
    
    if _, ok := p.Loc.(string); ok {
        return 0, 0, false
    }
    
    if locMap, ok := p.Loc.(map[string]interface{}); ok {
        if x, ok := locMap["x"].(float64); ok {
            if y, ok := locMap["y"].(float64); ok {
                return int(x), int(y), true
            }
        }
    }
    
    return 0, 0, false
}