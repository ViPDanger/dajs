package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/domain/repository"
)

type CharacterUsecase interface {
	New(ctx context.Context, character *entity.Character) (*string, error)
	Get(ctx context.Context, creator_id string, ids ...string) ([]*entity.Character, error)
	Set(ctx context.Context, character *entity.Character) error
	Delete(ctx context.Context, id string) error
}

func NewCharacterUseCase(Repository repository.CharacterRepository) CharacterUsecase {
	return &characterUsecase{CharacterRepository: Repository}
}

type characterUsecase struct {
	CharacterRepository repository.CharacterRepository
}

func (u *characterUsecase) New(ctx context.Context, item *entity.Character) (id *string, err error) {
	if u.CharacterRepository == nil || ctx == nil {
		return nil, errors.New("characterUsecase.New(): Nill pointer repository")

	}
	if id, err = u.CharacterRepository.Insert(ctx, item); err != nil {
		err = fmt.Errorf("characterUsecase.New()/%w", err)
	}
	return
}

func (u *characterUsecase) Get(ctx context.Context, creator_id string, ids ...string) (items []*entity.Character, err error) {
	if u.CharacterRepository == nil || ctx == nil {
		return nil, errors.New("characterUsecase.GetByID(): Nill pointer repository")

	}
	if items, err = u.CharacterRepository.Get(ctx, creator_id, ids...); err != nil {
		err = fmt.Errorf("characterUsecase.GetByID()/%w", err)
	}
	return
}

func (u *characterUsecase) Set(ctx context.Context, item *entity.Character) (err error) {
	if u.CharacterRepository == nil || ctx == nil {
		return errors.New("characterUsecase.Set(): Nill pointer repository")

	}
	if err = u.CharacterRepository.Update(ctx, item); err != nil {
		err = fmt.Errorf("characterUsecase.Update()/%w", err)
	}
	return
}
func (u *characterUsecase) Delete(ctx context.Context, id string) (err error) {
	if u.CharacterRepository == nil || ctx == nil {
		return errors.New("characterUsecase.Delete(): Nill pointer repository")

	}
	if err = u.CharacterRepository.Delete(ctx, id); err != nil {
		err = fmt.Errorf("characterUsecase.Delete()/%w", err)
	}
	return
}
