package models

import (
	"time"
)

type GameSummary struct {
	ID          uint64
	BoardSize   uint8
	Title       string
	Komi 		float32
	Date 		time.Time
	IsFinished	bool 	
	Rules 		Rules
	BlackPlayer *Player
	WhitePlayer *Player
}

type Game struct {
	GameSummary
	GameData    []GameState
}

type Player struct {
	Name string
	Rank uint8
}

type GameState struct {
	LastMove *Move
	Position Position
}

type Move struct {
	Point 	*Point
	IsPass  bool
	Color 	Color
}

type Position [][]PointState
type PointState uint8
type Color bool
type Rules uint8

type Point struct {
	X uint8
	Y uint8
}

const (
	Japanese Rules = iota
	Chinese
)

const (
	EmptyPoint PointState = iota
	BlackStone
	WhiteStone
)

const (
	Black Color = true
	White Color = false
)