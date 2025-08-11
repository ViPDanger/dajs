package repository

import (
	"context"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
)

type CharacterRepository interface {
	Insert(ctx context.Context, item *entity.Character) (*string, error)
	Get(ctx context.Context, creator_id string, ids ...string) ([]*entity.Character, error)
	Update(ctx context.Context, item *entity.Character) error
	Delete(ctx context.Context, id ...string) error
}
