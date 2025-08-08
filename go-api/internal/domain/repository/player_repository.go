package repository

import (
	"context"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
)

type PlayerCharRepository interface {
	Insert(ctx context.Context, item *entity.PlayerChar) (*string, error)
	Get(ctx context.Context, creator_id string, ids ...string) ([]*entity.PlayerChar, error)
	Update(ctx context.Context, item *entity.PlayerChar) error
	Delete(ctx context.Context, id string) error
}
