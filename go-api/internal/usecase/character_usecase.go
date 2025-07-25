package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/domain/repository"
)

type CharacterUsecase interface {
	New(ctx context.Context, character *entity.Character) (*entity.ID, error)
	GetByCreatorID(ctx context.Context, id entity.ID) ([]*entity.Character, error)
	GetByID(ctx context.Context, id entity.ID) (*entity.Character, error)
	GetArray(ctx context.Context, ids []entity.ID) ([]*entity.Character, error)
	GetAll(ctx context.Context) ([]*entity.Character, error)
	Set(ctx context.Context, character *entity.Character) error
	Delete(ctx context.Context, id entity.ID) error
}

func NewCharacterUseCase(Repository repository.CharacterRepository) CharacterUsecase {
	return &characterUsecase{CharacterRepository: Repository}
}

type characterUsecase struct {
	repository.CharacterRepository
}

func (u *characterUsecase) New(ctx context.Context, item *entity.Character) (id *entity.ID, err error) {
	if u.CharacterRepository == nil || ctx == nil {
		return nil, errors.New("characterUsecase.New(): Nill pointer repository")

	}
	if id, err = u.CharacterRepository.Insert(ctx, item); err != nil {
		err = fmt.Errorf("characterUsecase.New()/%w", err)
	}
	return
}
func (u *characterUsecase) GetByID(ctx context.Context, id entity.ID) (item *entity.Character, err error) {
	if u.CharacterRepository == nil || ctx == nil {
		return nil, errors.New("characterUsecase.GetByID(): Nill pointer repository")

	}
	if item, err = u.CharacterRepository.GetByID(ctx, id); err != nil {
		err = fmt.Errorf("characterUsecase.GetByID()/%w", err)
	}
	return
}
func (u *characterUsecase) GetArray(ctx context.Context, ids []entity.ID) (items []*entity.Character, err error) {
	if u.CharacterRepository == nil || ctx == nil {
		return nil, errors.New("characterUsecase.GetArray(): Nill pointer repository")

	}
	if items, err = u.CharacterRepository.GetArray(ctx, ids); err != nil {
		err = fmt.Errorf("characterUsecase.GetArray()/%w", err)
	}
	return
}
func (u *characterUsecase) GetAll(ctx context.Context) (items []*entity.Character, err error) {
	if u.CharacterRepository == nil || ctx == nil {
		return nil, errors.New("characterUsecase.GetAll(): Nill pointer repository")

	}
	if items, err = u.CharacterRepository.GetAll(ctx); err != nil {
		err = fmt.Errorf("characterUsecase.GetAll()/%w", err)
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
func (u *characterUsecase) Delete(ctx context.Context, id entity.ID) (err error) {
	if u.CharacterRepository == nil || ctx == nil {
		return errors.New("characterUsecase.Delete(): Nill pointer repository")

	}
	if err = u.CharacterRepository.Delete(ctx, id); err != nil {
		err = fmt.Errorf("characterUsecase.Delete()/%w", err)
	}
	return
}
