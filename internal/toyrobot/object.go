package toyrobot

type Object interface {
	Name() string

	container() container

	Place(container, Transform) error
	Transform() (Transform, error)
}

type MovableObject interface {
	Object

	Move() error
	RotateLeft() error
	RotateRight() error
}
