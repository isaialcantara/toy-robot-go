package toyrobot

type MovableObject interface {
	Object

	move() error
	rotateLeft() error
	rotateRight() error
}

func Move(movableObject MovableObject) error {
	return movableObject.move()
}

func RotateLeft(movableObject MovableObject) error {
	return movableObject.rotateLeft()
}

func RotateRight(movableObject MovableObject) error {
	return movableObject.rotateRight()
}

type movableObject struct {
	object
}

func (m *movableObject) move() error {
	if m.container() == nil {
		return NilContainerError
	}

	return moveObject(m.container(), m)
}

func (m *movableObject) rotateLeft() error {
	if m.container() == nil {
		return NilContainerError
	}

	return rotateObjectLeft(m.container(), m)
}

func (m *movableObject) rotateRight() error {
	if m.container() == nil {
		return NilContainerError
	}

	return rotateObjectRight(m.container(), m)
}
