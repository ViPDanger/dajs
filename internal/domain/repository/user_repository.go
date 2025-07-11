package repository

import (
	"context"

	"github.com/ViPDanger/dajs/internal/domain/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, item *entity.User) (*entity.ID, error)
	GetByID(ctx context.Context, id entity.ID) (*entity.User, error)
	GetArray(ctx context.Context, ids []entity.ID) ([]*entity.User, error)
	GetAll(ctx context.Context) ([]*entity.User, error)
	Update(ctx context.Context, item *entity.User) error
	Delete(ctx context.Context, id entity.ID) error
}
