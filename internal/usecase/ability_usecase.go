package usecase

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
)

type abilityUsecase struct {
	UseCase[entity.Ability]
}

func NewAbilityUseCase(Repository repository.Repository[entity.Ability]) UseCase[entity.Ability] {
	return &abilityUsecase{UseCase: NewDefaultUsecase(Repository)}
}
