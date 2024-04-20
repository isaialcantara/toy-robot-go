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
