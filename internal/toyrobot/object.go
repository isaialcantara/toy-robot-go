package toyrobot

type Object interface {
	Name() string

	Container() Container
	setContainer(Container)
	Place(Container, Transform) error
	Remove() error
	Transform() (Transform, error)
}
