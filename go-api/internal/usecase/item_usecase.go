package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/domain/repository"
)

type ItemUsecase interface {
	New(ctx context.Context, Item *entity.Item) (*entity.ID, error)
	GetByID(ctx context.Context, id entity.ID) (*entity.Item, error)
	GetArray(ctx context.Context, ids []entity.ID) ([]*entity.Item, error)
	GetAll(ctx context.Context) ([]*entity.Item, error)
	Set(ctx context.Context, Item *entity.Item) error
	Delete(ctx context.Context, id entity.ID) error
}

func NewItemUseCase(Repository repository.ItemRepository) ItemUsecase {
	return &itemUsecase{ItemRepository: Repository}
}

type itemUsecase struct {
	repository.ItemRepository
}

func (u *itemUsecase) New(ctx context.Context, item *entity.Item) (id *entity.ID, err error) {
	if u.ItemRepository == nil {
		return nil, errors.New("ItemRepository.New(): Nill pointer repository")

	}
	if id, err = u.ItemRepository.Insert(ctx, item); err != nil {
		err = fmt.Errorf("itemUsecase.New()/%w", err)
	}
	return
}
func (u *itemUsecase) GetByID(ctx context.Context, id entity.ID) (item *entity.Item, err error) {
	if u.ItemRepository == nil {
		return nil, errors.New("itemUsecase.GetByID(): Nill pointer repository")

	}
	if item, err = u.ItemRepository.GetByID(ctx, id); err != nil {
		err = fmt.Errorf("itemUsecase.GetByID()/%w", err)
	}
	return
}
func (u *itemUsecase) GetArray(ctx context.Context, ids []entity.ID) (items []*entity.Item, err error) {
	if u.ItemRepository == nil {
		return nil, errors.New("itemUsecase.GetArray(): Nill pointer repository")

	}
	if items, err = u.ItemRepository.GetArray(ctx, ids); err != nil {
		err = fmt.Errorf("itemUsecase.GetArray()/%w", err)
	}
	return
}
func (u *itemUsecase) GetAll(ctx context.Context) (items []*entity.Item, err error) {
	if u.ItemRepository == nil {
		return nil, errors.New("itemUsecase.GetAll(): Nill pointer repository")

	}
	if items, err = u.ItemRepository.GetAll(ctx); err != nil {
		err = fmt.Errorf("itemUsecase.GetAll()/%w", err)
	}
	return
}
func (u *itemUsecase) Set(ctx context.Context, item *entity.Item) (err error) {
	if u.ItemRepository == nil {
		return errors.New("itemUsecase.Update(): Nill pointer repository")

	}
	if err = u.ItemRepository.Update(ctx, item); err != nil {
		err = fmt.Errorf("itemUsecase.Update()/%w", err)
	}
	return
}
func (u *itemUsecase) Delete(ctx context.Context, id entity.ID) (err error) {
	if u.ItemRepository == nil {
		return errors.New("itemUsecase.Delete(): Nill pointer repository")

	}
	if err = u.ItemRepository.Delete(ctx, id); err != nil {
		err = fmt.Errorf("itemUsecase.Delete()/%w", err)
	}
	return
}
