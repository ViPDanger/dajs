package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/domain/repository"
)

type PlayerUsecase interface {
	New(ctx context.Context, Player *entity.PlayerCharacter) (*string, error)
	Get(ctx context.Context, id string) ([]*entity.PlayerCharacter, error)
	Set(ctx context.Context, Player *entity.PlayerCharacter) error
	Delete(ctx context.Context, id string) error
}

func NewPlayerCharacterUsecase(Repository repository.PlayerCharacterRepository) PlayerUsecase {
	return &playerUsecase{PlayerCharacterRepository: Repository}
}

type playerUsecase struct {
	repository.PlayerCharacterRepository
}

func (u *playerUsecase) New(ctx context.Context, item *entity.PlayerCharacter) (id *string, err error) {
	if u.PlayerCharacterRepository == nil {
		return nil, errors.New("playerUsecase.New(): Nill pointer repository")

	}
	if id, err = u.PlayerCharacterRepository.Insert(ctx, item); err != nil {
		err = fmt.Errorf("playerUsecase.New()/%w", err)
	}
	return
}
func (u *playerUsecase) Get(ctx context.Context, id string) (item []*entity.PlayerCharacter, err error) {
	if u.PlayerCharacterRepository == nil {
		return nil, errors.New("playerUsecase.GetByID(): Nill pointer repository")
	}
	if item, err = u.PlayerCharacterRepository.Get(ctx, id); err != nil {
		err = fmt.Errorf("playerUsecase.GetByID()/%w", err)
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
func (u *playerUsecase) Delete(ctx context.Context, id string) (err error) {
	if u.PlayerCharacterRepository == nil {
		return errors.New("playerUsecase.Delete(): Nill pointer repository")

	}
	if err = u.PlayerCharacterRepository.Delete(ctx, id); err != nil {
		err = fmt.Errorf("playerUsecase.Delete()/%w", err)
	}
	return
}
