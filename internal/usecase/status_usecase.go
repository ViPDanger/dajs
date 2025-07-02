package usecase

import (
	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/domain/repository"
)

type statusUsecase struct {
	UseCase[entity.Status]
}

func NewStatusUseCase(Repository repository.Repository[entity.Status]) UseCase[entity.Status] {
	return &statusUsecase{UseCase: NewDefaultUsecase(Repository)}
}
