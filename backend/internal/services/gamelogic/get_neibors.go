package gamelogic

import "github.com/J0hnLenin/Tesuji/internal/models"

func getNeibors(pos models.Position, p models.Point) []models.Point {
	neibors := make([]models.Point, 0, 4)
	size := len(pos)

	x := int(p.X)
	y := int(p.Y)

	if x-1 >= 0 {
		neibors = append(neibors, *newPoint(p.X-1, p.Y))
	}
	if x+1 < size {
		neibors = append(neibors, *newPoint(p.X+1, p.Y))
	}
	if y-1 >= 0 {
		neibors = append(neibors, *newPoint(p.X, p.Y-1))
	}
	if y+1 < size {
		neibors = append(neibors, *newPoint(p.X, p.Y+1))
	}
	return neibors
}