package usecase

import (
	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/domain/repository"
)

type itemUseCase struct {
	UseCase[entity.Item]
}

func NewItemUseCase(Repository repository.Repository[entity.Item]) UseCase[entity.Item] {
	return &itemUseCase{UseCase: NewDefaultUsecase(Repository)}
}
