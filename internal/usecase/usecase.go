package usecase

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
)

type UseCase[T entity.Identifiable] interface {
	New(item *T) error
	GetByID(id string) (*T, error)
	GetAll() ([]T, error)
	Set(item *T) error
	Delete(id string) error
}

type DefaultUseCase[T entity.Identifiable] struct {
	Repo repository.Repository[T]
}

func (uc *DefaultUseCase[T]) New(object *T) error {
	err := uc.Repo.Insert(object)
	return err

}
func (uc *DefaultUseCase[T]) Set(object *T) error {
	return uc.Repo.Update(object)

}
func (uc *DefaultUseCase[T]) GetByID(id string) (*T, error) {
	return uc.Repo.GetByID(id)
}
func (uc *DefaultUseCase[T]) GetAll() ([]T, error) {
	return uc.Repo.GetAll()
}
func (uc *DefaultUseCase[T]) Delete(id string) (err error) {
	return uc.Repo.Delete(id)
}
