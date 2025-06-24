package usecase

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
)

type GlossaryUsecase struct {
	UseCase[entity.Glossary]
}

func NewGlossaryUseCase(repository repository.Repository[entity.Glossary]) UseCase[entity.Glossary] {
	return &GlossaryUsecase{UseCase: NewDefaultUsecase(repository)}
}
