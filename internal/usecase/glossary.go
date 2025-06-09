package usecase

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
)

type GlossaryUseCase struct {
	GlossaryRepository repository.Repository[entity.Glossary]
}

func (uc *GlossaryUseCase) New(glossary *entity.Glossary) error {
	return uc.GlossaryRepository.New(glossary.ID, glossary)
}
func (uc *GlossaryUseCase) Set(glossary *entity.Glossary) error {
	return uc.GlossaryRepository.Set(glossary.ID, glossary)
}
func (uc *GlossaryUseCase) GetByID(id string) (r *entity.Glossary, err error) {
	return uc.GlossaryRepository.GetByID(id)
}
func (uc *GlossaryUseCase) GetAll() (r *[]entity.Glossary, err error) {
	return uc.GlossaryRepository.GetAll()
}
func (uc *GlossaryUseCase) Delete(id string) (err error) {
	return uc.GlossaryRepository.Delete(id)
}
