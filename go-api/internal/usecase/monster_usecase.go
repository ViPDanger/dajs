package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/domain/repository"
)

type MonsterUsecase interface {
	New(ctx context.Context, Monster *entity.Monster) (*entity.ID, error)
	GetByID(ctx context.Context, id entity.ID) (*entity.Monster, error)
	GetByCreatorID(ctx context.Context, id entity.ID) ([]*entity.Monster, error)
	GetArray(ctx context.Context, ids []entity.ID) ([]*entity.Monster, error)
	GetAll(ctx context.Context) ([]*entity.Monster, error)
	Set(ctx context.Context, Monster *entity.Monster) error
	Delete(ctx context.Context, id entity.ID) error
}

func NewMonsterUsecase(Repository repository.MonsterRepository) MonsterUsecase {
	return &monsterUsecase{MonsterRepository: Repository}
}

type monsterUsecase struct {
	repository.MonsterRepository
}

func (u *monsterUsecase) New(ctx context.Context, item *entity.Monster) (id *entity.ID, err error) {
	if u.MonsterRepository == nil {
		return nil, errors.New("monsterUsecase.New(): Nill pointer repository")

	}
	if id, err = u.MonsterRepository.Insert(ctx, item); err != nil {
		err = fmt.Errorf("monsterUsecase.New()/%w", err)
	}
	return
}
func (u *monsterUsecase) GetByID(ctx context.Context, id entity.ID) (item *entity.Monster, err error) {
	if u.MonsterRepository == nil {
		return nil, errors.New("monsterUsecase.GetByID(): Nill pointer repository")

	}
	if item, err = u.MonsterRepository.GetByID(ctx, id); err != nil {
		err = fmt.Errorf("monsterUsecase.GetByID()/%w", err)
	}
	return
}
func (u *monsterUsecase) GetArray(ctx context.Context, ids []entity.ID) (items []*entity.Monster, err error) {
	if u.MonsterRepository == nil {
		return nil, errors.New("monsterUsecase.GetArray(): Nill pointer repository")

	}
	if items, err = u.MonsterRepository.GetArray(ctx, ids); err != nil {
		err = fmt.Errorf("monsterUsecase.GetArray()/%w", err)
	}
	return
}
func (u *monsterUsecase) GetAll(ctx context.Context) (items []*entity.Monster, err error) {
	if u.MonsterRepository == nil {
		return nil, errors.New("monsterUsecase.GetAll(): Nill pointer repository")

	}
	if items, err = u.MonsterRepository.GetAll(ctx); err != nil {
		err = fmt.Errorf("monsterUsecase.GetAll()/%w", err)
	}
	return
}
func (u *monsterUsecase) Set(ctx context.Context, item *entity.Monster) (err error) {
	if u.MonsterRepository == nil {
		return errors.New("monsterUsecase.Set(): Nill pointer repository")

	}
	if err = u.MonsterRepository.Update(ctx, item); err != nil {
		err = fmt.Errorf("monsterUsecase.Update()/%w", err)
	}
	return
}
func (u *monsterUsecase) Delete(ctx context.Context, id entity.ID) (err error) {
	if u.MonsterRepository == nil {
		return errors.New("monsterUsecase.Delete(): Nill pointer repository")

	}
	if err = u.MonsterRepository.Delete(ctx, id); err != nil {
		err = fmt.Errorf("monsterUsecase.Delete()/%w", err)
	}
	return
}
