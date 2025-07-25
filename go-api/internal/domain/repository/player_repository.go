package repository

import (
	"context"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
)

type PlayerCharacterRepository interface {
	Insert(ctx context.Context, item *entity.PlayerCharacter) (*entity.ID, error)
	GetByCreatorID(ctx context.Context, id entity.ID) ([]*entity.PlayerCharacter, error)
	GetByID(ctx context.Context, id entity.ID) (*entity.PlayerCharacter, error)
	GetArray(ctx context.Context, ids []entity.ID) ([]*entity.PlayerCharacter, error)
	GetAll(ctx context.Context) ([]*entity.PlayerCharacter, error)
	Update(ctx context.Context, item *entity.PlayerCharacter) error
	Delete(ctx context.Context, id entity.ID) error
}
