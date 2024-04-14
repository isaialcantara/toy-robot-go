package toyrobot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Robot_Place(t *testing.T) {
	t.Run("fail when container is nil", func(t *testing.T) {
		robot := NewRobot("C3PO")
		err := robot.Place(nil, NewTransform(1, 1, North))
		assert.ErrorIs(t, err, NilContainerError)
	})
	t.Run("place a robot on a table", func(t *testing.T) {
		table := NewTable(5, 5)

		robot := NewRobot("R2D2")
		transform := NewTransform(1, 1, South)

		err := robot.Place(&table, transform)
		if assert.NoError(t, err) {
			assert.Equal(t, &table, robot.container())
			returnedTransform, err := table.objectTransform(&robot)
			assert.NoError(t, err)
			assert.Equal(t, transform, returnedTransform)
		}
	})

	t.Run("place the robot on the same table", func(t *testing.T) {
		table := NewTable(5, 5)
		robot := NewRobot("R2D2")
		transform := NewTransform(1, 1, South)

		{
			err := robot.Place(&table, transform)
			assert.NoError(t, err)
		}

		otherTransform := NewTransform(2, 0, East)
		if err := robot.Place(&table, otherTransform); assert.NoError(t, err) {
			assert.Equal(t, &table, robot.container())
			returnedTransform, err := table.objectTransform(&robot)
			assert.NoError(t, err)
			assert.Equal(t, otherTransform, returnedTransform)
		}
	})

	t.Run("place the robot on another table", func(t *testing.T) {
		table := NewTable(5, 5)
		robot := NewRobot("R2D2")
		transform := NewTransform(1, 1, South)

		{
			err := robot.Place(&table, transform)
			assert.NoError(t, err)
		}

		otherTable := NewTable(3, 3)
		otherTransform := NewTransform(1, 2, West)
		if err := robot.Place(&otherTable, otherTransform); assert.NoError(t, err) {
			assert.Equal(t, &otherTable, robot.container())
			returnedTransform, err := otherTable.objectTransform(&robot)
			assert.NoError(t, err)
			assert.Equal(t, otherTransform, returnedTransform)

			{
				_, err := table.objectTransform(&robot)
				assert.ErrorIs(t, err, NotPlacedError)
			}
		}
	})
}
