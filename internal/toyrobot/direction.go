package toyrobot

type Direction string

const (
	North Direction = "North"
	East  Direction = "East"
	South Direction = "South"
	West  Direction = "West"
)

func (d Direction) String() string { return string(d) }

func AllDirections() [4]Direction {
	return [4]Direction{North, East, South, West}
}

func (d *Direction) RotateRight() {
	switch *d {
	case North:
		*d = East
	case East:
		*d = South
	case South:
		*d = West
	case West:
		*d = North
	}
}

func (d *Direction) RotateLeft() {
	switch *d {
	case North:
		*d = West
	case East:
		*d = North
	case South:
		*d = East
	case West:
		*d = South
	}
}

func (d Direction) Step(initial Vec2) Vec2 {
	final := initial

	switch d {
	case North:
		final.Y++
	case East:
		final.X++
	case South:
		final.Y--
	case West:
		final.X--
	}

	return final
}
