package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/domain/repository"
)

type NPCUsecase interface {
	New(ctx context.Context, NPC *entity.NPC) (*entity.ID, error)
	GetByID(ctx context.Context, id entity.ID) (*entity.NPC, error)
	GetByCreatorID(ctx context.Context, id entity.ID) ([]*entity.NPC, error)
	GetArray(ctx context.Context, ids []entity.ID) ([]*entity.NPC, error)
	GetAll(ctx context.Context) ([]*entity.NPC, error)
	Set(ctx context.Context, NPC *entity.NPC) error
	Delete(ctx context.Context, id entity.ID) error
}

func NewNPCUsecase(Repository repository.NPCRepository) NPCUsecase {
	return &npcUsecase{NPCRepository: Repository}
}

type npcUsecase struct {
	repository.NPCRepository
}

func (u *npcUsecase) New(ctx context.Context, item *entity.NPC) (id *entity.ID, err error) {
	if u.NPCRepository == nil {
		return nil, errors.New("npcUsecase.New(): Nill pointer repository")

	}
	if id, err = u.NPCRepository.Insert(ctx, item); err != nil {
		err = fmt.Errorf("npcUsecase.New()/%w", err)
	}
	return
}
func (u *npcUsecase) GetByID(ctx context.Context, id entity.ID) (item *entity.NPC, err error) {
	if u.NPCRepository == nil {
		return nil, errors.New("npcUsecase.GetByID(): Nill pointer repository")

	}
	if item, err = u.NPCRepository.GetByID(ctx, id); err != nil {
		err = fmt.Errorf("npcUsecase.GetByID()/%w", err)
	}
	return
}
func (u *npcUsecase) GetArray(ctx context.Context, ids []entity.ID) (items []*entity.NPC, err error) {
	if u.NPCRepository == nil {
		return nil, errors.New("npcUsecase.GetArray(): Nill pointer repository")

	}
	if items, err = u.NPCRepository.GetArray(ctx, ids); err != nil {
		err = fmt.Errorf("npcUsecase.GetArray()/%w", err)
	}
	return
}
func (u *npcUsecase) GetAll(ctx context.Context) (items []*entity.NPC, err error) {
	if u.NPCRepository == nil {
		return nil, errors.New("npcUsecase.GetAll(): Nill pointer repository")

	}
	if items, err = u.NPCRepository.GetAll(ctx); err != nil {
		err = fmt.Errorf("npcUsecase.GetAll()/%w", err)
	}
	return
}
func (u *npcUsecase) Set(ctx context.Context, item *entity.NPC) (err error) {
	if u.NPCRepository == nil {
		return errors.New("npcUsecase.Set(): Nill pointer repository")

	}
	if err = u.NPCRepository.Update(ctx, item); err != nil {
		err = fmt.Errorf("npcUsecase.Update()/%w", err)
	}
	return
}
func (u *npcUsecase) Delete(ctx context.Context, id entity.ID) (err error) {
	if u.NPCRepository == nil {
		return errors.New("npcUsecase.Delete(): Nill pointer repository")

	}
	if err = u.NPCRepository.Delete(ctx, id); err != nil {
		err = fmt.Errorf("npcUsecase.Delete(ctx context.Context,)/%w", err)
	}
	return
}
