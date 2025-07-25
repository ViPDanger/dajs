package repository

import (
	"context"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
)

type PlayerCharRepository interface {
	Insert(ctx context.Context, item *entity.PlayerChar) (*entity.ID, error)
	GetByCreatorID(ctx context.Context, id entity.ID) ([]entity.PlayerChar, error)
	GetByID(ctx context.Context, id entity.ID) (*entity.PlayerChar, error)
	GetArray(ctx context.Context, ids []entity.ID) ([]entity.PlayerChar, error)
	GetAll(ctx context.Context) ([]entity.PlayerChar, error)
	Update(ctx context.Context, item *entity.PlayerChar) error
	Delete(ctx context.Context, id entity.ID) error
}
