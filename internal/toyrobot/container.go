package toyrobot

type container interface {
	placeObject(Object, Transform) error
	placeObjectOnOther(Object, container, Transform) error
	objectTransform(Object) (Transform, error)
	moveObject(MovableObject) error
	rotateObjectLeft(MovableObject) error
	rotateObjectRight(MovableObject) error
}
