package matchmaking

import "errors"

const (
	boardSize = 9
)

var (
	errInvalidPoisition = errors.New("")
)

type Side uint

const (
	Blue Side = iota
	Red
)

type BoardPosition uint

const (
	TopLeft BoardPosition = iota
	TopCenter
	TopRight
	CenterLeft
	Center
	CenterRight
	BottomLeft
	BottomCenter
	BottomRight
)

type Card struct {
	name  string
	up    int
	left  int
	right int
	down  int
}

type Board struct {
	cardBoard [9]*Card
	turn      Side
}

func NewBoard() *Board {
	return &Board{
		turn: Blue,
	}
}

func (board *Board) Play(card Card, playerColor Side, where int) error {
	//check this is ok
	if where < 0 || where > boardSize || board.cardBoard[where] != nil {
		return errInvalidPoisition
	}

}

type Player struct {
	id   string
	name string
	side Side
	hand []Card
}
