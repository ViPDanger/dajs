package repository

import (
	"context"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
)

type PlayerCharacterRepository interface {
	Insert(ctx context.Context, item *entity.PlayerCharacter) (*string, error)
	Get(ctx context.Context, creator_id string, ids ...string) ([]*entity.PlayerCharacter, error)
	Update(ctx context.Context, item *entity.PlayerCharacter) error
	Delete(ctx context.Context, id string) error
}
