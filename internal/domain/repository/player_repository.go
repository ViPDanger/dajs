package repository

import (
	"context"

	"github.com/ViPDanger/dajs/internal/domain/entity"
)

type PlayerRepository interface {
	Insert(ctx context.Context, item *entity.Player) (*entity.ID, error)
	GetByID(ctx context.Context, id entity.ID) (*entity.Player, error)
	GetArray(ctx context.Context, ids []entity.ID) ([]*entity.Player, error)
	GetAll(ctx context.Context) ([]*entity.Player, error)
	Update(ctx context.Context, item *entity.Player) error
	Delete(ctx context.Context, id entity.ID) error
}
