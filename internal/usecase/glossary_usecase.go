package usecase

import (
	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/domain/repository"
)

type GlossaryUsecase struct {
	UseCase[entity.Glossary]
}

func NewGlossaryUseCase(repository repository.Repository[entity.Glossary]) UseCase[entity.Glossary] {
	return &GlossaryUsecase{UseCase: NewDefaultUsecase(repository)}
}
