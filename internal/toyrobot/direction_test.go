package toyrobot

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AllDirections(t *testing.T) {
	assert.Equal(t, [4]Direction{North, East, South, West}, AllDirections())
}

func Test_DirectionRotateLeft(t *testing.T) {
	leftRotations := map[Direction]Direction{
		North: West,
		East:  North,
		South: East,
		West:  South,
	}

	for initial, final := range leftRotations {
		name := fmt.Sprintf("rotate direction %s to the left, ending at %s", initial, final)
		t.Run(name, func(t *testing.T) {
			initial.RotateLeft()
			assert.Equal(t, final, initial)
		})
	}
}

func Test_DirectionRotateRight(t *testing.T) {
	rightRotations := map[Direction]Direction{
		North: East,
		East:  South,
		South: West,
		West:  North,
	}

	for initial, final := range rightRotations {
		name := fmt.Sprintf("rotate direction %s to the right, ending at %s", initial, final)
		t.Run(name, func(t *testing.T) {
			initial.RotateRight()
			assert.Equal(t, final, initial)
		})
	}
}

func Test_DirectionStep(t *testing.T) {
	directionToFinalPosition := map[Direction]Position{
		North: {0, 1},
		East:  {1, 0},
		South: {0, -1},
		West:  {-1, 0},
	}

	for direction, finalPosition := range directionToFinalPosition {
		name := fmt.Sprintf("move position according to direction %s", direction)
		t.Run(name, func(t *testing.T) {
			initialPosition := Position{0, 0}
			assert.Equal(t, finalPosition, direction.Step(initialPosition))
		})
	}
}
