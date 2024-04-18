package toyrobot

type Obstacle struct {
	Object
}

func NewObstacle() Obstacle {
	return Obstacle{&object{}}
}
