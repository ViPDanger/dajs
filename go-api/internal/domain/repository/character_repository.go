package repository

import (
	"context"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
)

type CharacterRepository interface {
	Insert(ctx context.Context, item *entity.Character) (*entity.ID, error)
	GetByID(ctx context.Context, id entity.ID) (*entity.Character, error)
	GetByCreatorID(ctx context.Context, id entity.ID) ([]*entity.Character, error)
	GetArray(ctx context.Context, ids []entity.ID) ([]*entity.Character, error)
	GetAll(ctx context.Context) ([]*entity.Character, error)
	Update(ctx context.Context, item *entity.Character) error
	Delete(ctx context.Context, id entity.ID) error
}
