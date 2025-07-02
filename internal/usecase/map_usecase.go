package usecase

import (
	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/domain/repository"
)

type mapUsecase struct {
	UseCase[entity.Map]
}

func NewMapUseCase(Repository repository.Repository[entity.Map]) UseCase[entity.Map] {
	return &mapUsecase{UseCase: NewDefaultUsecase(Repository)}
}
