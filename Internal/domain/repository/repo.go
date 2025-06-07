package repository

type Repository[T any] interface {
	Save(item *T) error
	GetByID(id string) (*T, error)
}
