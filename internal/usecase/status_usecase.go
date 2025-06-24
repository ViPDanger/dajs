package usecase

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
)

type statusUsecase struct {
	UseCase[entity.Status]
}

func NewStatusUseCase(Repository repository.Repository[entity.Status]) UseCase[entity.Status] {
	return &statusUsecase{UseCase: NewDefaultUsecase(Repository)}
}
