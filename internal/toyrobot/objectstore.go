package toyrobot

import (
	mapset "github.com/deckarep/golang-set/v2"
)

const (
	NilObjectError            = constError("the object cannot be nil")
	PositionAlreadyTakenError = constError("the position has already been taken")
	ObjectNotFoundError       = constError("the object wasn't found in the store")
)

type objectStore interface {
	addObject(Object, Transform) error
	removeObject(Object) error
	objectTransform(Object) (Transform, error)
}

type store struct {
	transforms     map[Object]Transform
	takenPositions mapset.Set[Position]
}

func newStore() store {
	return store{
		transforms:     make(map[Object]Transform),
		takenPositions: mapset.NewSet[Position](),
	}
}

func (s *store) addObject(object Object, transform Transform) error {
	if object == nil {
		return NilObjectError
	}

	if s.takenPositions.Contains(transform.Position) {
		return PositionAlreadyTakenError
	}

	if currentTransform, objectAlreadyStored := s.transforms[object]; objectAlreadyStored {
		s.takenPositions.Remove(currentTransform.Position)
	}

	s.transforms[object] = transform
	s.takenPositions.Add(transform.Position)

	return nil
}

func (s *store) removeObject(object Object) error {
	if object == nil {
		return NilObjectError
	}

	if currentTransform, objectAlreadyStored := s.transforms[object]; objectAlreadyStored {
		s.takenPositions.Remove(currentTransform.Position)
		delete(s.transforms, object)
		return nil
	}

	return ObjectNotFoundError
}

func (s *store) objectTransform(object Object) (Transform, error) {
	if object == nil {
		return Transform{}, NilObjectError
	}

	if transform, objectAlreadyStored := s.transforms[object]; objectAlreadyStored {
		return transform, nil
	}

	return Transform{}, ObjectNotFoundError
}
