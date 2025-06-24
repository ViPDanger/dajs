package usecase

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
)

type itemUseCase struct {
	UseCase[entity.Item]
}

func NewItemUseCase(Repository repository.Repository[entity.Item]) UseCase[entity.Item] {
	return &itemUseCase{UseCase: NewDefaultUsecase(Repository)}
}
