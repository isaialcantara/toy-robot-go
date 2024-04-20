package toyrobot

import (
	"math"
)

type CircleTable struct {
	objectStore

	Radius int
}

func NewCircleTable(radius int) CircleTable {
	store := newStore()

	return CircleTable{
		objectStore: &store,
		Radius:      radius,
	}
}

func (c *CircleTable) inBounds(pos Position) bool {
	distance := math.Sqrt(float64(pos.X*pos.X) + float64(float64(pos.Y*pos.Y)))

	return distance <= float64(c.Radius)
}
