package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/domain/repository"
)

type PlayerUsecase interface {
	New(ctx context.Context, Player *entity.Player) (*entity.ID, error)
	GetByID(ctx context.Context, id entity.ID) (*entity.Player, error)
	GetArray(ctx context.Context, ids []entity.ID) ([]*entity.Player, error)
	GetAll(ctx context.Context) ([]*entity.Player, error)
	Set(ctx context.Context, Player *entity.Player) error
	Delete(ctx context.Context, id entity.ID) error
}

func NewPlayerUsecase(Repository repository.PlayerRepository) PlayerUsecase {
	return &playerUsecase{PlayerRepository: Repository}
}

type playerUsecase struct {
	repository.PlayerRepository
}

func (u *playerUsecase) New(ctx context.Context, item *entity.Player) (id *entity.ID, err error) {
	if u.PlayerRepository == nil {
		return nil, errors.New("playerUsecase.New(): Nill pointer repository")

	}
	if id, err = u.PlayerRepository.Insert(ctx, item); err != nil {
		err = fmt.Errorf("playerUsecase.New()/%w", err)
	}
	return
}
func (u *playerUsecase) GetByID(ctx context.Context, id entity.ID) (item *entity.Player, err error) {
	if u.PlayerRepository == nil {
		return nil, errors.New("playerUsecase.GetByID(): Nill pointer repository")

	}
	if item, err = u.PlayerRepository.GetByID(ctx, id); err != nil {
		err = fmt.Errorf("playerUsecase.GetByID()/%w", err)
	}
	return
}
func (u *playerUsecase) GetArray(ctx context.Context, ids []entity.ID) (items []*entity.Player, err error) {
	if u.PlayerRepository == nil {
		return nil, errors.New("playerUsecase.GetArray(): Nill pointer repository")

	}
	if items, err = u.PlayerRepository.GetArray(ctx, ids); err != nil {
		err = fmt.Errorf("playerUsecase.GetArray()/%w", err)
	}
	return
}
func (u *playerUsecase) GetAll(ctx context.Context) (items []*entity.Player, err error) {
	if u.PlayerRepository == nil {
		return nil, errors.New("playerUsecase.GetAll(): Nill pointer repository")

	}
	if items, err = u.PlayerRepository.GetAll(ctx); err != nil {
		err = fmt.Errorf("playerUsecase.GetAll()/%w", err)
	}
	return
}
func (u *playerUsecase) Set(ctx context.Context, item *entity.Player) (err error) {
	if u.PlayerRepository == nil {
		return errors.New("playerUsecase.Set(): Nill pointer repository")

	}
	if err = u.PlayerRepository.Update(ctx, item); err != nil {
		err = fmt.Errorf("playerUsecase.Update()/%w", err)
	}
	return
}
func (u *playerUsecase) Delete(ctx context.Context, id entity.ID) (err error) {
	if u.PlayerRepository == nil {
		return errors.New("playerUsecase.Delete(): Nill pointer repository")

	}
	if err = u.PlayerRepository.Delete(ctx, id); err != nil {
		err = fmt.Errorf("playerUsecase.Delete()/%w", err)
	}
	return
}
