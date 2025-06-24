package usecase

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
)

type MapUsecase struct {
	UseCase[entity.Map]
}

func NewMapUseCase(Repository repository.Repository[entity.Map]) UseCase[entity.Map] {
	return &MapUsecase{UseCase: NewDefaultUsecase(Repository)}
}
