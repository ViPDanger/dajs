package usecase

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
)

type GlossaryUseCase struct {
	defaultUseCase[entity.Glossary]
}

func NewGlossaryUseCase(repository repository.Repository[entity.Glossary]) UseCase[entity.Glossary] {
	return NewDefaultUsecase(repository)
}
