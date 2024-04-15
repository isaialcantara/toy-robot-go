package toyrobot

import (
	"log"

	mapset "github.com/deckarep/golang-set/v2"
)

const (
	NilObjectError            = constError("the object cannot be nil")
	NilContainerError         = constError("the container cannot be nil")
	NotPlacedError            = constError("the object hasn't been placed on the table")
	InvalidPlacementError     = constError("the object cannot be placed outside of the table")
	PositionAlreadyTakenError = constError("the position has already been taken")
	InvalidMovementError      = constError("the object cannot be moved to fall from the table")
)

type Table struct {
	Width          int
	Length         int
	transforms     map[Object]Transform
	takenPositions mapset.Set[Vec2]
}

func NewTable(width int, length int) Table {
	return Table{
		Width:          width,
		Length:         length,
		transforms:     make(map[Object]Transform),
		takenPositions: mapset.NewSet[Vec2](),
	}
}

func (t *Table) placeObject(object Object, transform Transform) error {
	if object == nil {
		return NilObjectError
	}

	if err := t.addObject(object, transform); err != nil {
		return err
	}

	if object.Container() == nil {
		object.setContainer(t)
		return nil
	}

	if object.Container() != t {
		if err := object.Container().removeObject(object); err != nil {
			log.Println("object container reference desync")
		}

		object.setContainer(t)
	}

	return nil
}

func (t *Table) addObject(object Object, transform Transform) error {
	if object == nil {
		return NilObjectError
	}

	if !t.contains(transform.Position) {
		return InvalidPlacementError
	}

	currentTransform, objectAlreadyOnTable := t.transforms[object]

	if objectAlreadyOnTable && currentTransform.Position == transform.Position {
		t.transforms[object] = transform
		return nil
	}

	if t.takenPositions.Contains(transform.Position) {
		return PositionAlreadyTakenError
	}

	if objectAlreadyOnTable {
		t.takenPositions.Remove(currentTransform.Position)
	}

	t.takenPositions.Add(transform.Position)
	t.transforms[object] = transform

	return nil
}

func (t *Table) removeObject(object Object) error {
	if transform, objectAlreadyOnTable := t.transforms[object]; objectAlreadyOnTable {
		delete(t.transforms, object)
		t.takenPositions.Remove(transform.Position)
		return nil
	}

	return NotPlacedError
}

func (t Table) objectTransform(object Object) (Transform, error) {
	if object == nil {
		return Transform{}, NilObjectError
	}

	if transform, objectAlreadyOnTable := t.transforms[object]; objectAlreadyOnTable {
		return transform, nil
	}

	return Transform{}, NotPlacedError
}

func (t *Table) moveObject(object MovableObject) error {
	if object == nil {
		return NilObjectError
	}

	currentTransform, objectAlreadyOnTable := t.transforms[object]
	if !objectAlreadyOnTable {
		return NotPlacedError
	}

	newPosition := currentTransform.Direction.Step(currentTransform.Position)

	if !t.contains(newPosition) {
		return InvalidMovementError
	}

	if t.takenPositions.Contains(newPosition) {
		return PositionAlreadyTakenError
	}

	newTransform := currentTransform
	newTransform.Position = newPosition

	t.transforms[object] = newTransform
	t.takenPositions.Remove(currentTransform.Position)
	t.takenPositions.Add(newPosition)

	return nil
}

func (t *Table) rotateObjectLeft(object MovableObject) error {
	if object == nil {
		return NilObjectError
	}
	transform, objectAlreadyOnTable := t.transforms[object]
	if !objectAlreadyOnTable {
		return NotPlacedError
	}

	transform.Direction.RotateLeft()
	return nil
}

func (t *Table) rotateObjectRight(object MovableObject) error {
	if object == nil {
		return NilObjectError
	}

	transform, objectAlreadyOnTable := t.transforms[object]
	if !objectAlreadyOnTable {
		return NotPlacedError
	}

	transform.Direction.RotateRight()
	return nil
}

func (t Table) contains(pos Vec2) bool {
	return (pos.X >= 0 && pos.X < t.Width) && (pos.Y >= 0 && pos.Y < t.Length)
}
