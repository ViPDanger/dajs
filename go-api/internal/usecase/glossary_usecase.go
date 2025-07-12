package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/domain/repository"
)

type GlossaryUsecase interface {
	New(ctx context.Context, object *entity.Glossary) (*entity.ID, error)
	GetByID(ctx context.Context, id entity.ID) (*entity.Glossary, error)
	GetArray(ctx context.Context, ids []entity.ID) ([]*entity.Glossary, error)
	GetAll(ctx context.Context) ([]*entity.Glossary, error)
	Set(ctx context.Context, object *entity.Glossary) error
	Delete(ctx context.Context, id entity.ID) error
}

func NewGlossaryUseCase(repository repository.GlossaryRepository) GlossaryUsecase {
	return &glossaryUsecase{GlossaryRepository: repository}
}

type glossaryUsecase struct {
	repository.GlossaryRepository
}

func (u *glossaryUsecase) New(ctx context.Context, item *entity.Glossary) (id *entity.ID, err error) {
	if u.GlossaryRepository == nil {
		return nil, errors.New("glossaryUsecase.New(): Nill pointer repository")

	}
	if id, err = u.GlossaryRepository.Insert(ctx, item); err != nil {
		err = fmt.Errorf("glossaryUsecase.New()/%w", err)
	}
	return
}
func (u *glossaryUsecase) GetByID(ctx context.Context, id entity.ID) (item *entity.Glossary, err error) {
	if u.GlossaryRepository == nil {
		return nil, errors.New("glossaryUsecase.GetByID(): Nill pointer repository")

	}
	if item, err = u.GlossaryRepository.GetByID(ctx, id); err != nil {
		err = fmt.Errorf("glossaryUsecase.GetByID()/%w", err)
	}
	return
}
func (u *glossaryUsecase) GetArray(ctx context.Context, ids []entity.ID) (items []*entity.Glossary, err error) {
	if u.GlossaryRepository == nil {
		return nil, errors.New("glossaryUsecase.GetArray(): Nill pointer repository")

	}
	if items, err = u.GlossaryRepository.GetArray(ctx, ids); err != nil {
		err = fmt.Errorf("glossaryUsecase.GetArray()/%w", err)
	}
	return
}
func (u *glossaryUsecase) GetAll(ctx context.Context) (items []*entity.Glossary, err error) {
	if u.GlossaryRepository == nil {
		return nil, errors.New("glossaryUsecase.GetAll(): Nill pointer repository")

	}
	if items, err = u.GlossaryRepository.GetAll(ctx); err != nil {
		err = fmt.Errorf("glossaryUsecase.GetAll()/%w", err)
	}
	return
}
func (u *glossaryUsecase) Set(ctx context.Context, item *entity.Glossary) (err error) {
	if u.GlossaryRepository == nil {
		return errors.New("glossaryUsecase.Set(): Nill pointer repository")

	}
	if err = u.GlossaryRepository.Update(ctx, item); err != nil {
		err = fmt.Errorf("glossaryUsecase.Update()/%w", err)
	}
	return
}
func (u *glossaryUsecase) Delete(ctx context.Context, id entity.ID) (err error) {
	if u.GlossaryRepository == nil {
		return errors.New("glossaryUsecase.Delete(): Nill pointer repository")

	}
	if err = u.GlossaryRepository.Delete(ctx, id); err != nil {
		err = fmt.Errorf("glossaryUsecase.Delete()/%w", err)
	}
	return
}
