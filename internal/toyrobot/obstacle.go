package toyrobot

const ObstacleNotPlacedError = constError("the robot hasn't been placed yet")

type Obstacle struct {
	name      string
	container container
}

func NewObstacle(name string) Obstacle {
	return Obstacle{name: name}
}

func (o Obstacle) Name() string {
	return o.name
}

func (o Obstacle) Container() container {
	return o.container
}

func (o *Obstacle) setContainer(container container) {
	o.container = container
}

func (o *Obstacle) Place(container container, transform Transform) error {
	if container == nil {
		return NilContainerError
	}

	return container.placeObject(o, transform)
}

func (o *Obstacle) Transform() (Transform, error) {
	if o.Container() == nil {
		return Transform{}, ObstacleNotPlacedError
	}

	return o.Container().objectTransform(o)
}
