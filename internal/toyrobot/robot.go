package toyrobot

const RobotNotPlacedError = constError("the robot hasn't been placed yet")

type Robot struct {
	name      string
	container container
}

func NewRobot(name string) Robot {
	return Robot{name: name}
}

func (r Robot) Name() string {
	return r.name
}

func (r Robot) Container() container {
	return r.container
}

func (r *Robot) setContainer(container container) {
	r.container = container
}

func (r *Robot) Place(container container, transform Transform) error {
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
