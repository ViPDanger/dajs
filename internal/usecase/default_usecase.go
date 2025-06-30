package usecase

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
	"fmt"
	"reflect"
)

type UseCase[T entity.Identifiable] interface {
	New(item *T) error
	GetByID(id string) (*T, error)
	GetArray(ids []string) ([]T, error)
	GetAll() ([]T, error)
	Set(item *T) error
	Delete(id string) error
}

type defaultUseCase[T entity.Identifiable] struct {
	Repo repository.Repository[T]
}

func NewDefaultUsecase[T entity.Identifiable](repository repository.Repository[T]) UseCase[T] {
	return &defaultUseCase[T]{
		Repo: repository,
	}
}
func (uc *defaultUseCase[T]) New(object *T) error {
	if object == nil {
		TName := reflect.TypeOf(new(T)).String()[1:]
		return fmt.Errorf("defaultUseCase[%s].New(): Nill pointer", TName)
	}
	if err := uc.Repo.Insert(object); err != nil {
		TName := reflect.TypeOf(new(T)).String()[1:]
		return fmt.Errorf("defaultUseCase[%s].New()/%w", TName, err)
	}

	return nil

}
func (uc *defaultUseCase[T]) Set(object *T) error {
	TName := reflect.TypeOf(new(T)).Name()[1:]
	if object == nil {
		return fmt.Errorf("defaultUseCase[%s].Set(): Nill pointer", TName)
	}
	if err := uc.Repo.Update(object); err != nil {
		return fmt.Errorf("defaultUseCase[%s].Set()/%w", TName, err)
	}
	return nil

}
func (uc *defaultUseCase[T]) GetByID(id string) (*T, error) {
	if id == "" {
		TName := reflect.TypeOf(new(T)).Name()[1:]
		return nil, fmt.Errorf("defaultUseCase[%s].Get(): Nill pointer", TName)
	}
	ret, err := uc.Repo.GetByID(id)
	if err != nil {
		TName := reflect.TypeOf(new(T)).Name()[1:]
		return nil, fmt.Errorf("defaultUseCase[%s].Set()/%w", TName, err)
	}
	return ret, nil
}

func (uc *defaultUseCase[T]) GetArray(ids []string) (ret []T, err error) {

	ret, err = uc.Repo.GetArray(ids)
	if err != nil {
		TName := reflect.TypeOf(new(T)).Name()[1:]
		return nil, fmt.Errorf("defaultUseCase[%s].GetArray()/%w", TName, err)
	}
	return
}

func (uc *defaultUseCase[T]) GetAll() (ret []T, err error) {
	ret, err = uc.Repo.GetAll()
	if err != nil {
		TName := reflect.TypeOf(*new(T)).String()
		return nil, fmt.Errorf("defaultUseCase[%s].GetAll()/%w", TName, err)
	}
	return
}
func (uc *defaultUseCase[T]) Delete(id string) (err error) {
	if err = uc.Repo.Delete(id); err != nil {
		TName := reflect.TypeOf(new(T)).Name()[1:]
		return fmt.Errorf("defaultUseCase[%s].GetAll()/%w", TName, err)
	}
	return
}
