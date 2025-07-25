package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/domain/repository"
)

type PlayerUsecase interface {
	New(ctx context.Context, Player *entity.PlayerCharacter) (*entity.ID, error)
	GetByID(ctx context.Context, id entity.ID) (*entity.PlayerCharacter, error)
	GetByCreatorID(ctx context.Context, id entity.ID) ([]*entity.PlayerCharacter, error)
	GetArray(ctx context.Context, ids []entity.ID) ([]*entity.PlayerCharacter, error)
	GetAll(ctx context.Context) ([]*entity.PlayerCharacter, error)
	Set(ctx context.Context, Player *entity.PlayerCharacter) error
	Delete(ctx context.Context, id entity.ID) error
}

func NewPlayerCharacterUsecase(Repository repository.PlayerCharacterRepository) PlayerUsecase {
	return &playerUsecase{PlayerCharacterRepository: Repository}
}

type playerUsecase struct {
	repository.PlayerCharacterRepository
}

func (u *playerUsecase) New(ctx context.Context, item *entity.PlayerCharacter) (id *entity.ID, err error) {
	if u.PlayerCharacterRepository == nil {
		return nil, errors.New("playerUsecase.New(): Nill pointer repository")

	}
	if id, err = u.PlayerCharacterRepository.Insert(ctx, item); err != nil {
		err = fmt.Errorf("playerUsecase.New()/%w", err)
	}
	return
}
func (u *playerUsecase) GetByID(ctx context.Context, id entity.ID) (item *entity.PlayerCharacter, err error) {
	if u.PlayerCharacterRepository == nil {
		return nil, errors.New("playerUsecase.GetByID(): Nill pointer repository")

	}
	if item, err = u.PlayerCharacterRepository.GetByID(ctx, id); err != nil {
		err = fmt.Errorf("playerUsecase.GetByID()/%w", err)
	}
	return
}
func (u *playerUsecase) GetArray(ctx context.Context, ids []entity.ID) (items []*entity.PlayerCharacter, err error) {
	if u.PlayerCharacterRepository == nil {
		return nil, errors.New("playerUsecase.GetArray(): Nill pointer repository")

	}
	if items, err = u.PlayerCharacterRepository.GetArray(ctx, ids); err != nil {
		err = fmt.Errorf("playerUsecase.GetArray()/%w", err)
	}
	return
}
func (u *playerUsecase) GetAll(ctx context.Context) (items []*entity.PlayerCharacter, err error) {
	if u.PlayerCharacterRepository == nil {
		return nil, errors.New("playerUsecase.GetAll(): Nill pointer repository")

	}
	if items, err = u.PlayerCharacterRepository.GetAll(ctx); err != nil {
		err = fmt.Errorf("playerUsecase.GetAll()/%w", err)
	}
	return
}
func (u *playerUsecase) Set(ctx context.Context, item *entity.PlayerCharacter) (err error) {
	if u.PlayerCharacterRepository == nil {
		return errors.New("playerUsecase.Set(): Nill pointer repository")

	}
	if err = u.PlayerCharacterRepository.Update(ctx, item); err != nil {
		err = fmt.Errorf("playerUsecase.Update()/%w", err)
	}
	return
}
func (u *playerUsecase) Delete(ctx context.Context, id entity.ID) (err error) {
	if u.PlayerCharacterRepository == nil {
		return errors.New("playerUsecase.Delete(): Nill pointer repository")

	}
	if err = u.PlayerCharacterRepository.Delete(ctx, id); err != nil {
		err = fmt.Errorf("playerUsecase.Delete()/%w", err)
	}
	return
}
