package toyrobot

type Object interface {
	Name() string

	Container() container
	setContainer(container)
	Place(container, Transform) error
	Remove() error
	Transform() (Transform, error)
}
