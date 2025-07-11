package repository

import (
	"context"

	"github.com/ViPDanger/dajs/internal/domain/entity"
)

type GlossaryRepository interface {
	Insert(ctx context.Context, item *entity.Glossary) (*entity.ID, error)
	GetByID(ctx context.Context, id entity.ID) (*entity.Glossary, error)
	GetArray(ctx context.Context, ids []entity.ID) ([]*entity.Glossary, error)
	GetAll(ctx context.Context) ([]*entity.Glossary, error)
	Update(ctx context.Context, item *entity.Glossary) error
	Delete(ctx context.Context, id entity.ID) error
}
