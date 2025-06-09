package repository

import (
	"errors"
)

type Repository[T any] interface {
	New(id string, item *T) error
	GetByID(id string) (*T, error)
	GetAll() (*[]T, error)
	Set(id string, item *T) error
	Delete(id string) error
}

type TestRepository[T any] struct {
	Objects map[string]T
}

func NewTestRepository[T any]() *TestRepository[T] {
	return &TestRepository[T]{Objects: make(map[string]T)}
}
func (r *TestRepository[T]) New(id string, item *T) error {
	if _, exists := r.Objects[id]; exists {
		return errors.New("Object exists")
	}
	r.Objects[id] = *item
	return nil
}

func (r *TestRepository[T]) GetByID(id string) (ret *T, err error) {
	if object, exists := r.Objects[id]; !exists {
		err = errors.New("Object not exists")
	} else {
		ret = &object
	}
	return
}
func (r *TestRepository[T]) GetAll() (ret *[]T, err error) {
	m := make([]T, len(r.Objects))
	var i int
	for _, object := range r.Objects {
		m[i] = object
	}
	ret = &m
	return
}
func (r *TestRepository[T]) Set(id string, item *T) error {
	if _, exists := r.Objects[id]; !exists {
		return errors.New("Object not exists")
	}
	r.Objects[id] = *item
	return nil

}
func (r *TestRepository[T]) Delete(id string) (err error) {
	if _, exists := r.Objects[id]; !exists {
		err = errors.New("Object not exists")
	}
	delete(r.Objects, id)
	return
}
