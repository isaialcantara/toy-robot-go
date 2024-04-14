package toyrobot

const RobotNotPlacedError = constError("the robot hasn't been placed yet")

type Robot struct {
	name           string
	robotContainer container
}

func NewRobot(name string) Robot {
	return Robot{name: name}
}

func (r Robot) Name() string {
	return r.name
}

func (r Robot) container() container {
	return r.robotContainer
}

func (r *Robot) Place(container container, transform Transform) error {
	if container == nil {
		return NilContainerError
	}

	if r.container() == nil {
		if err := container.placeObject(r, transform); err != nil {
			return err
		}

		r.robotContainer = container
		return nil
	}

	if r.container() != container {
		if err := r.container().placeObjectOnOther(r, container, transform); err != nil {
			return err
		}

		r.robotContainer = container
		return nil
	}

	if err := r.container().placeObject(r, transform); err != nil {
		return err
	}

	return nil
}

func (r *Robot) Transform() (Transform, error) {
	if r.container() == nil {
		return Transform{}, RobotNotPlacedError
	}

	return r.container().objectTransform(r)
}

func (r *Robot) Move() error {
	if r.container() == nil {
		return RobotNotPlacedError
	}

	return r.container().moveObject(r)
}

func (r *Robot) RotateLeft() error {
	if r.container() == nil {
		return RobotNotPlacedError
	}

	return r.container().rotateObjectLeft(r)
}

func (r *Robot) RotateRight() error {
	if r.container() == nil {
		return RobotNotPlacedError
	}

	return r.container().rotateObjectRight(r)
}
