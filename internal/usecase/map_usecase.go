package usecase

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
)

type mapUsecase struct {
	UseCase[entity.Map]
}

func NewMapUseCase(Repository repository.Repository[entity.Map]) UseCase[entity.Map] {
	return &mapUsecase{UseCase: NewDefaultUsecase(Repository)}
}
