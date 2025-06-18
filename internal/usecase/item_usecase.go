package usecase

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
)

type itemUseCase struct {
	ArmorRepo repository.Repository[entity.Armor]
}

func NewItemUseCase(repository repository.Repository[entity.Item]) UseCase[entity.Item] {
	return NewDefaultUsecase(repository)
}
