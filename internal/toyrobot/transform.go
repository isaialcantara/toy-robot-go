package toyrobot

type Position struct {
	X int
	Y int
}

type Transform struct {
	Position  Position
	Direction Direction
}

func NewTransform(x int, y int, direction Direction) Transform {
	return Transform{
		Position:  Position{x, y},
		Direction: direction,
	}
}
