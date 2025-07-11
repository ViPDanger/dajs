package repository

import (
	"context"

	"github.com/ViPDanger/dajs/internal/domain/entity"
)

type MonsterRepository interface {
	Insert(ctx context.Context, item *entity.Monster) (*entity.ID, error)
	GetByID(ctx context.Context, id entity.ID) (*entity.Monster, error)
	GetArray(ctx context.Context, ids []entity.ID) ([]*entity.Monster, error)
	GetAll(ctx context.Context) ([]*entity.Monster, error)
	Update(ctx context.Context, item *entity.Monster) error
	Delete(ctx context.Context, id entity.ID) error
}
