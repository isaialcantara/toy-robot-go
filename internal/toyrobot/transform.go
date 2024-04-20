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

func (t Transform) Step() Transform {
	switch t.Direction {
	case North:
		t.Position.Y++
	case East:
		t.Position.X++
	case South:
		t.Position.Y--
	case West:
		t.Position.X--
	}

	return t
}

func (t Transform) RotateRight() Transform {
	switch t.Direction {
	case North:
		t.Direction = East
	case East:
		t.Direction = South
	case South:
		t.Direction = West
	case West:
		t.Direction = North
	}

	return t
}

func (t Transform) RotateLeft() Transform {
	switch t.Direction {
	case North:
		t.Direction = West
	case East:
		t.Direction = North
	case South:
		t.Direction = East
	case West:
		t.Direction = South
	}

	return t
}
