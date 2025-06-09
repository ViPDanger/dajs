package repository

type Repository[T any] interface {
	New(item *T) error
	GetByID(id string) (*T, error)
	GetAll() (*[]T, error)
	Set(item *T) error
	Delete(id string) error
}
