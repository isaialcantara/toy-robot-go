package toyrobot

type Object interface {
	container() Container
	setContainer(Container)
	place(Container, Transform) error
	remove() error
	transform() (Transform, error)
}

func GetContainer(object Object) Container {
	return object.container()
}

func Place(object Object, container Container, transform Transform) error {
	return object.place(container, transform)
}

func Remove(object Object) error {
	return object.remove()
}

func GetTransform(object Object) (Transform, error) {
	return object.transform()
}

type object struct {
	myContainer Container
}

func (o *object) container() Container {
	return o.myContainer
}

func (o *object) setContainer(container Container) {
	o.myContainer = container
}

func (o *object) place(container Container, transform Transform) error {
	if container == nil {
		return NilContainerError
	}

	return placeObject(container, o, transform)
}

func (o *object) remove() error {
	if o.container() == nil {
		return NilContainerError
	}

	return removeObject(o.container(), o)
}

func (o *object) transform() (Transform, error) {
	if o.container() == nil {
		return Transform{}, NilContainerError
	}

	return objectTransform(o.container(), o)
}
