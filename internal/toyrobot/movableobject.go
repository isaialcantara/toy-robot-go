package toyrobot

type MovableObject interface {
	Object

	Move() error
	RotateLeft() error
	RotateRight() error
}
