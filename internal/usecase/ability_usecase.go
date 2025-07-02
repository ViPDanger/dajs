package usecase

import (
	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/domain/repository"
)

type abilityUsecase struct {
	UseCase[entity.Ability]
}

func NewAbilityUseCase(Repository repository.Repository[entity.Ability]) UseCase[entity.Ability] {
	return &abilityUsecase{UseCase: NewDefaultUsecase(Repository)}
}
