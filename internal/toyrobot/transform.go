package toyrobot

type Vec2 struct {
	X int
	Y int
}

type Transform struct {
	Position  Vec2
	Direction Direction
}

func NewTransform(x int, y int, direction Direction) Transform {
	return Transform{
		Position:  Vec2{x, y},
		Direction: direction,
	}
}
