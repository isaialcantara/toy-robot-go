package toyrobot

const RobotNotPlacedError = constError("the robot hasn't been placed yet")

type Robot struct {
	name      string
	container Container
}

func NewRobot(name string) Robot {
	return Robot{name: name}
}

func (r Robot) Name() string {
	return r.name
}

func (r Robot) Container() Container {
	return r.container
}

func (r *Robot) setContainer(container Container) {
	r.container = container
}

func (r *Robot) Place(container Container, transform Transform) error {
	if container == nil {
		return NilContainerError
	}

	return container.placeObject(r, transform)
}

func (r *Robot) Remove() error {
	if r.Container() == nil {
		return RobotNotPlacedError
	}

	return r.Container().removeObject(r)
}

func (r *Robot) Transform() (Transform, error) {
	if r.Container() == nil {
		return Transform{}, RobotNotPlacedError
	}

	return r.Container().objectTransform(r)
}

func (r *Robot) Move() error {
	if r.Container() == nil {
		return RobotNotPlacedError
	}

	return r.Container().moveObject(r)
}

func (r *Robot) RotateLeft() error {
	if r.Container() == nil {
		return RobotNotPlacedError
	}

	return r.Container().rotateObjectLeft(r)
}

func (r *Robot) RotateRight() error {
	if r.Container() == nil {
		return RobotNotPlacedError
	}

	return r.Container().rotateObjectRight(r)
}
