package repository

import "github.com/ViPDanger/dajs/internal/domain/entity"

type Repository[T entity.Identifiable] interface {
	Insert(item *T) error
	GetByID(id string) (*T, error)
	GetArray(ids []string) ([]T, error)
	GetAll() ([]T, error)
	Update(item *T) error
	Delete(id string) error
}
