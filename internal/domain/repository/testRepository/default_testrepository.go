package testrepository

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/ViPDanger/dajs/internal/domain/entity"
)

func NewTestRepository[T entity.Identifiable]() *testRepository[T] {
	return &testRepository[T]{Objects: make(map[string]T)}
}
func (r *testRepository[T]) GetArray(ids []string) (ret []T, err error) {
	ret = make([]T, len(ids))
	for i := range ids {
		object, err := r.GetByID(ids[i])
		if err != nil {
			TName := reflect.TypeOf(ret[0]).Name()
			return nil, fmt.Errorf("Test Repositrory[%s].GetArray/%w", TName, err)
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
	var o T
	TName := reflect.TypeOf(o).Name()
	if item == nil {
		return fmt.Errorf("Test Repositrory[%s].Insert: Nill pointer", TName)
	}
	if _, exists := r.Objects[(*item).GetID()]; exists {
		return fmt.Errorf("Test Repositrory[%s].Insert: Object already exists", TName)
	}
	r.Objects[(*item).GetID()] = *item
	return nil
}

func (r *testRepository[T]) Update(item *T) error {

	if item == nil {
		var o T
		TName := reflect.TypeOf(o).Name()
		return fmt.Errorf("Test Repositrory[%s].Update: Nill pointer", TName)
	}
	if _, exists := r.Objects[(*item).GetID()]; !exists {
		var o T
		TName := reflect.TypeOf(o).Name()
		return fmt.Errorf("Test Repositrory[%s].Update: Object not exists", TName)
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
		var o T
		TName := reflect.TypeOf(o).Name()
		err = fmt.Errorf("Test Repositrory[%s].Delete: Object not exists", TName)
	}
	delete(r.Objects, id)
	return
}

type testRepository[T entity.Identifiable] struct {
	Objects map[string]T
}
