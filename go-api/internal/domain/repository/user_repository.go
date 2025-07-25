package repository

import (
	"context"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, item *entity.User) error
	Get(ctx context.Context, username string) (*entity.User, error)
	Update(ctx context.Context, item *entity.User) error
	Delete(ctx context.Context, id entity.ID) error
}
