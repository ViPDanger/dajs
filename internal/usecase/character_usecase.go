package usecase

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
)

type characterUseCase struct {
	defaultUseCase[entity.Character]
}

func NewCharacterUseCase(repository repository.Repository[entity.Character]) UseCase[entity.Character] {
	return NewDefaultUsecase(repository)
}
