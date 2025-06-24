package usecase

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
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
	err := uc.Repo.Insert(object)
	return err

}
func (uc *defaultUseCase[T]) Set(object *T) error {
	return uc.Repo.Update(object)

}
func (uc *defaultUseCase[T]) GetByID(id string) (*T, error) {
	return uc.Repo.GetByID(id)
}

func (uc *defaultUseCase[T]) GetArray(ids []string) ([]T, error) {
	return uc.Repo.GetArray(ids)
}
func (uc *defaultUseCase[T]) GetAll() ([]T, error) {
	return uc.Repo.GetAll()
}
func (uc *defaultUseCase[T]) Delete(id string) (err error) {
	return uc.Repo.Delete(id)
}
