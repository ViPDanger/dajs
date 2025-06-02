package repository

type Repository[T any] interface {
	Save(item *T) error
	GetByID(id int64) (*T, error)
}
