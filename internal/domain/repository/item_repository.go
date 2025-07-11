package repository

import (
	"context"

	"github.com/ViPDanger/dajs/internal/domain/entity"
)

type ItemRepository interface {
	Insert(ctx context.Context, item *entity.Item) (*entity.ID, error)
	GetByID(ctx context.Context, id entity.ID) (*entity.Item, error)
	GetArray(ctx context.Context, ids []entity.ID) ([]*entity.Item, error)
	GetAll(ctx context.Context) ([]*entity.Item, error)
	Update(ctx context.Context, item *entity.Item) error
	Delete(ctx context.Context, id entity.ID) error
}
