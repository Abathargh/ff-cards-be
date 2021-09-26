package matchmaking

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
	cardBoard []Card
	turn      Side
}

func NewBoard() *Board {
	return &Board{
		turn: Blue,
	}
}

type Player struct {
	id   string
	name string
	side Side
	hand []Card
}
