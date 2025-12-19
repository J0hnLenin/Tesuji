package kgsservice

import "encoding/json"

func parseGames(jsonData string) ([]game, error) {
	var response response
	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		return nil, err
	}

	var games []game
	for _, message := range response.Messages {
		if len(message.Games) > 0 {
			for _, game := range message.Games {
				if game.GameType != "challenge" {
					games = append(games, game)
				}
			}
		}
	}

	return games, nil
}