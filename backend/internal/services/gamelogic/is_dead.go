package gamelogic

import "github.com/J0hnLenin/Tesuji/internal/models"

func isDead(pos models.Position, p models.Point) (bool, map[models.Point]bool) {
	visited := map[models.Point]bool{}
	toVisit := []models.Point{p}
	pointState := pos[p.X][p.Y]
	if pointState == models.EmptyPoint {
		return false, nil
	}
	for len(toVisit) > 0 {
		toVisit, p = toVisit[:len(toVisit)-1], toVisit[len(toVisit)-1]
		neibors := getNeibors(pos, p)
		for _, n := range neibors {
			if pos[n.X][n.Y] == models.EmptyPoint {
				return false, nil
			}
			if (!visited[n]) && (pos[n.X][n.Y] == pointState) {
				toVisit = append(toVisit, n)
			}
		}
		visited[p] = true
	}
	return true, visited
}
