package testrepository

import (
	"DAJ/internal/domain/entity"
	"errors"
)

func NewTestRepository[T entity.Identifiable]() *testRepository[T] {
	return &testRepository[T]{Objects: make(map[string]T)}
}
func (r *testRepository[T]) GetArray(ids []string) (ret []T, err error) {
	ret = make([]T, len(ids))
	for i := range ids {
		object, err := r.GetByID(ids[i])
		if err != nil {
			return nil, err
		}
		ret[i] = *object
	}
	return
}

func (r *testRepository[T]) GetByID(id string) (ret *T, err error) {
	if object, exists := r.Objects[id]; !exists {
		err = errors.New("Object not exists")
	} else {
		ret = &object
	}
	return
}
func (r *testRepository[T]) Insert(item *T) error {

	if _, exists := r.Objects[(*item).GetID()]; exists {
		return errors.New("Object exists")
	}
	r.Objects[(*item).GetID()] = *item
	return nil
}

func (r *testRepository[T]) Update(item *T) error {
	if _, exists := r.Objects[(*item).GetID()]; !exists {
		return errors.New("Object not exists")
	}
	r.Objects[(*item).GetID()] = *item
	return nil

}
func (r *testRepository[T]) GetAll() (ret []T, err error) {
	ret = make([]T, len(r.Objects))
	var i int
	for _, object := range r.Objects {
		ret[i] = object
	}
	return
}
func (r *testRepository[T]) Delete(id string) (err error) {
	if _, exists := r.Objects[id]; !exists {
		err = errors.New("Object not exists")
	}
	delete(r.Objects, id)
	return
}

type testRepository[T entity.Identifiable] struct {
	Objects map[string]T
}
