package toyrobot

type container interface {
	placeObject(Object, Transform) error
	removeObject(Object) error
	objectTransform(Object) (Transform, error)
	moveObject(MovableObject) error
	rotateObjectLeft(MovableObject) error
	rotateObjectRight(MovableObject) error
}
