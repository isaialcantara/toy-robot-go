package toyrobot

const (
	NilContainerError     = constError("the container cannot be nil")
	ObjectNotPlacedError  = constError("the object hasn't been placed yet")
	InvalidPlacementError = constError("the object cannot be placed out of bounds")
	InvalidMovementError  = constError("the object cannot be moved out of bounds")
)

type Container interface {
	objectStore

	inBounds(Position) bool
}

func placeObject(container Container, object Object, transform Transform) error {
	if container == nil {
		return NilContainerError
	}

	if object == nil {
		return NilObjectError
	}

	if !container.inBounds(transform.Position) {
		return InvalidPlacementError
	}

	if err := container.addObject(object, transform); err != nil {
		return err
	}

	if object.container() == nil {
		object.setContainer(container)
		return nil
	}

	if object.container() != container {
		object.container().removeObject(object)
		object.setContainer(container)
	}

	return nil
}

func removeObject(container Container, object Object) error {
	return container.removeObject(object)
}

func objectTransform(container Container, object Object) (Transform, error) {
	return container.objectTransform(object)
}

func moveObject(container Container, object MovableObject) error {
	if container == nil {
		return NilContainerError
	}

	if object == nil {
		return NilObjectError
	}

	if currentTransform, err := container.objectTransform(object); err == nil {
		nextTransform := currentTransform.Step()

		if container.inBounds(nextTransform.Position) {
			return container.addObject(object, nextTransform)
		}

		return InvalidMovementError
	}

	return ObjectNotPlacedError
}

func rotateObjectLeft(container Container, object MovableObject) error {
	if container == nil {
		return NilContainerError
	}

	if object == nil {
		return NilObjectError
	}

	if currentTransform, err := container.objectTransform(object); err == nil {
		newTransform := currentTransform.RotateLeft()
		return container.addObject(object, newTransform)
	}

	return ObjectNotPlacedError
}

func rotateObjectRight(container Container, object MovableObject) error {
	if container == nil {
		return NilContainerError
	}

	if object == nil {
		return NilObjectError
	}

	if currentTransform, err := container.objectTransform(object); err == nil {
		newTransform := currentTransform.RotateRight()
		return container.addObject(object, newTransform)
	}

	return ObjectNotPlacedError
}
