package testrepository

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
)

func NewTestRepository[T entity.Identifiable]() *testRepository[T] {
	return &testRepository[T]{Objects: make(map[string]T)}
}
func (r *testRepository[T]) GetArray(ctx context.Context, ids []entity.ID) (ret []*T, err error) {
	ret = make([]*T, len(ids))
	for i := range ids {
		object, err := r.GetByID(ctx, ids[i])
		if err != nil {
			TName := reflect.TypeOf(ret[0]).Name()
			return nil, fmt.Errorf("Test Repositrory[%s].GetArray/%w", TName, err)
		}
		ret[i] = object
	}
	return
}

func (r *testRepository[T]) GetByID(ctx context.Context, id entity.ID) (ret *T, err error) {
	if object, exists := r.Objects[id.String()]; !exists {
		err = errors.New("Object not exists")
	} else {
		ret = &object
	}
	return
}
func (r *testRepository[T]) Insert(ctx context.Context, item *T) (*entity.ID, error) {
	var o T
	TName := reflect.TypeOf(o).String()
	if item == nil {
		return nil, fmt.Errorf("Test Repositrory[%s].Insert: Nill pointer", TName)
	}
	if _, exists := r.Objects[(*item).GetID().String()]; exists {
		return nil, fmt.Errorf("Test Repositrory[%s].Insert: Object already exists", TName)
	}
	r.Objects[(*item).GetID().String()] = *item
	ret := (*item).GetID()
	return &ret, nil
}

func (r *testRepository[T]) Update(ctx context.Context, item *T) error {

	if item == nil {
		var o T
		TName := reflect.TypeOf(o).String()
		return fmt.Errorf("Test Repositrory[%s].Update: Nill pointer", TName)
	}
	if _, exists := r.Objects[(*item).GetID().String()]; !exists {
		var o T
		TName := reflect.TypeOf(o).String()
		return fmt.Errorf("Test Repositrory[%s].Update: Object not exists", TName)
	}
	r.Objects[(*item).GetID().String()] = *item
	return nil

}
func (r *testRepository[T]) GetAll(ctx context.Context) (ret []*T, err error) {
	ret = make([]*T, len(r.Objects))
	var i int
	for _, object := range r.Objects {
		ret[i] = &object
	}
	return
}
func (r *testRepository[T]) Delete(ctx context.Context, id entity.ID) (err error) {

	if _, exists := r.Objects[id.String()]; !exists {
		var o T
		TName := reflect.TypeOf(o).String()
		err = fmt.Errorf("Test Repositrory[%s].Delete: Object not exists", TName)
	}
	delete(r.Objects, id.String())
	return
}

type testRepository[T entity.Identifiable] struct {
	Objects map[string]T
}
