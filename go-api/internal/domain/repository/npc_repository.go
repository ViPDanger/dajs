package repository

import (
	"context"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
)

type NPCRepository interface {
	Insert(ctx context.Context, item *entity.NPC) (*entity.ID, error)
	GetByID(ctx context.Context, id entity.ID) (*entity.NPC, error)
	GetArray(ctx context.Context, ids []entity.ID) ([]*entity.NPC, error)
	GetAll(ctx context.Context) ([]*entity.NPC, error)
	Update(ctx context.Context, item *entity.NPC) error
	Delete(ctx context.Context, id entity.ID) error
}
