package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/domain/repository"
)

type PlayerCharUsecase interface {
	New(ctx context.Context, playerChar *entity.PlayerChar) (*entity.ID, error)
	GetByCreatorID(ctx context.Context, id entity.ID) ([]entity.PlayerChar, error)
	GetByID(ctx context.Context, id entity.ID) (*entity.PlayerChar, error)
	GetArray(ctx context.Context, ids []entity.ID) ([]entity.PlayerChar, error)
	GetAll(ctx context.Context) ([]entity.PlayerChar, error)
	Set(ctx context.Context, playerChar *entity.PlayerChar) error
	Delete(ctx context.Context, id entity.ID) error
}

func NewPlayerCharUsecase(Repository repository.PlayerCharRepository) PlayerCharUsecase {
	return &playerCharUsecase{PlayerCharRepository: Repository}
}

type playerCharUsecase struct {
	repository.PlayerCharRepository
}

func (u *playerCharUsecase) New(ctx context.Context, item *entity.PlayerChar) (id *entity.ID, err error) {
	if u.PlayerCharRepository == nil || ctx == nil {
		return nil, errors.New("playerCharUsecase.New(): Nill pointer repository")

	}
	if id, err = u.PlayerCharRepository.Insert(ctx, item); err != nil {
		err = fmt.Errorf("playerCharUsecase.New()/%w", err)
	}
	return
}
func (u *playerCharUsecase) GetByID(ctx context.Context, id entity.ID) (item *entity.PlayerChar, err error) {
	if u.PlayerCharRepository == nil || ctx == nil {
		return nil, errors.New("playerCharUsecase.GetByID(): Nill pointer repository")

	}
	if item, err = u.PlayerCharRepository.GetByID(ctx, id); err != nil {
		err = fmt.Errorf("playerCharUsecase.GetByID()/%w", err)
	}
	return
}
func (u *playerCharUsecase) GetByCreatorID(ctx context.Context, id entity.ID) (items []entity.PlayerChar, err error) {
	if u.PlayerCharRepository == nil || ctx == nil {
		return nil, errors.New("playerCharUsecase.GetByID(): Nill pointer repository")

	}
	if items, err = u.PlayerCharRepository.GetByCreatorID(ctx, id); err != nil {
		err = fmt.Errorf("playerCharUsecase.GetByID()/%w", err)
	}
	return
}
func (u *playerCharUsecase) GetArray(ctx context.Context, ids []entity.ID) (items []entity.PlayerChar, err error) {
	if u.PlayerCharRepository == nil || ctx == nil {
		return nil, errors.New("playerCharUsecase.GetArray(): Nill pointer repository")

	}
	if items, err = u.PlayerCharRepository.GetArray(ctx, ids); err != nil {
		err = fmt.Errorf("playerCharUsecase.GetArray()/%w", err)
	}
	return
}
func (u *playerCharUsecase) GetAll(ctx context.Context) (items []entity.PlayerChar, err error) {
	if u.PlayerCharRepository == nil || ctx == nil {
		return nil, errors.New("playerCharUsecase.GetAll(): Nill pointer repository")

	}
	if items, err = u.PlayerCharRepository.GetAll(ctx); err != nil {
		err = fmt.Errorf("playerCharUsecase.GetAll()/%w", err)
	}
	return
}
func (u *playerCharUsecase) Set(ctx context.Context, item *entity.PlayerChar) (err error) {
	if u.PlayerCharRepository == nil || ctx == nil {
		return errors.New("playerCharUsecase.Set(): Nill pointer repository")

	}
	if err = u.PlayerCharRepository.Update(ctx, item); err != nil {
		err = fmt.Errorf("playerCharUsecase.Update()/%w", err)
	}
	return
}
func (u *playerCharUsecase) Delete(ctx context.Context, id entity.ID) (err error) {
	if u.PlayerCharRepository == nil || ctx == nil {
		return errors.New("playerCharUsecase.Delete(): Nill pointer repository")

	}
	if err = u.PlayerCharRepository.Delete(ctx, id); err != nil {
		err = fmt.Errorf("playerCharUsecase.Delete()/%w", err)
	}
	return
}
