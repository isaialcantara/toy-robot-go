package toyrobot

type Object interface {
	Name() string

	Container() container
	setContainer(container)
	Place(container, Transform) error
	Remove() error
	Transform() (Transform, error)
}

type MovableObject interface {
	Object

	Move() error
	RotateLeft() error
	RotateRight() error
}
