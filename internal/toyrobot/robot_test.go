package toyrobot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PlaceRobot(t *testing.T) {
	t.Run("fail when container is nil", func(t *testing.T) {
		robot := NewRobot("C3PO")
		err := Place(&robot, nil, NewTransform(1, 1, North))
		assert.ErrorIs(t, err, NilContainerError)
	})
	t.Run("place a robot on a table", func(t *testing.T) {
		table := NewTable(5, 5)
		robot := NewRobot("R2D2")
		transform := NewTransform(1, 1, South)

		err := Place(robot, &table, transform)

		if assert.NoError(t, err) {
			assert.Equal(t, &table, GetContainer(robot))
			returnedTransform, err := GetTransform(robot)

			if assert.NoError(t, err) {
				assert.Equal(t, transform, returnedTransform)
			}
		}
	})

	t.Run("place the robot on the same table", func(t *testing.T) {
		table := NewTable(5, 5)
		robot := NewRobot("R2D2")
		transform := NewTransform(1, 1, South)

		err := Place(robot, &table, transform)

		if assert.NoError(t, err) {
			otherTransform := NewTransform(2, 0, East)

			if err := Place(robot, &table, otherTransform); assert.NoError(t, err) {
				assert.Equal(t, &table, GetContainer(robot))
				returnedTransform, err := GetTransform(robot)

				if assert.NoError(t, err) {
					assert.Equal(t, otherTransform, returnedTransform)
				}
			}
		}
	})

	t.Run("place the robot on another table", func(t *testing.T) {
		table := NewTable(5, 5)
		robot := NewRobot("R2D2")
		transform := NewTransform(1, 1, South)

		if err := Place(robot, &table, transform); assert.NoError(t, err) {
			otherTable := NewTable(3, 3)
			otherTransform := NewTransform(1, 2, West)

			if err := Place(robot, &otherTable, otherTransform); assert.NoError(t, err) {
				assert.Equal(t, &otherTable, GetContainer(robot))

				if returnedTransform, err := GetTransform(robot); assert.NoError(t, err) {
					assert.Equal(t, otherTransform, returnedTransform)
				}

				{
					_, err := table.objectTransform(robot.MovableObject)
					assert.ErrorIs(t, err, NotPlacedError)
				}

			}
		}
	})
}
