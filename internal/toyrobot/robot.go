package toyrobot

type Robot struct {
	MovableObject

	name string
}

func NewRobot(name string) Robot {
	return Robot{MovableObject: &movableObject{}, name: name}
}

func (r Robot) Name() string {
	return r.name
}
