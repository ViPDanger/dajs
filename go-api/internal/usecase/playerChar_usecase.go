package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/domain/repository"
)

type PlayerCharUsecase interface {
	New(ctx context.Context, character *entity.PlayerChar) (*string, error)
	Get(ctx context.Context, creator_id string, ids ...string) ([]*entity.PlayerChar, error)
	Set(ctx context.Context, character *entity.PlayerChar) error
	Delete(ctx context.Context, id string) error
}

func NewPlayerCharUsecase(Repository repository.PlayerCharRepository) PlayerCharUsecase {
	return &playerCharUsecase{PlayerCharRepository: Repository}
}

type playerCharUsecase struct {
	repository.PlayerCharRepository
}

func (u *playerCharUsecase) New(ctx context.Context, item *entity.PlayerChar) (id *string, err error) {
	if u.PlayerCharRepository == nil || ctx == nil {
		return nil, errors.New("playerCharUsecase.New(): Nill pointer repository")

	}
	if id, err = u.PlayerCharRepository.Insert(ctx, item); err != nil {
		err = fmt.Errorf("playerCharUsecase.New()/%w", err)
	}
	return
}

func (u *playerCharUsecase) Get(ctx context.Context, creator_id string, ids ...string) (items []*entity.PlayerChar, err error) {
	if u.PlayerCharRepository == nil || ctx == nil {
		return nil, errors.New("characterUsecase.GetByID(): Nill pointer repository")

	}
	if items, err = u.PlayerCharRepository.Get(ctx, creator_id, ids...); err != nil {
		err = fmt.Errorf("characterUsecase.GetByID()/%w", err)
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
func (u *playerCharUsecase) Delete(ctx context.Context, id string) (err error) {
	if u.PlayerCharRepository == nil || ctx == nil {
		return errors.New("playerCharUsecase.Delete(): Nill pointer repository")

	}
	if err = u.PlayerCharRepository.Delete(ctx, id); err != nil {
		err = fmt.Errorf("playerCharUsecase.Delete()/%w", err)
	}
	return
}
