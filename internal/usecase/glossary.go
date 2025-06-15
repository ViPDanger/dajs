package usecase

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
)

type GlossaryUseCase struct {
	Repo repository.Repository[entity.Glossary]
}

func (uc *GlossaryUseCase) New(glossary *entity.Glossary) error {
	return uc.Repo.Insert(glossary)
}
func (uc *GlossaryUseCase) Set(glossary *entity.Glossary) error {
	return uc.Repo.Update(glossary)
}
func (uc *GlossaryUseCase) GetByID(id string) (r *entity.Glossary, err error) {
	return uc.Repo.GetByID(id)
}
func (uc *GlossaryUseCase) GetAll() (r []entity.Glossary, err error) {
	return uc.Repo.GetAll()
}
func (uc *GlossaryUseCase) Delete(id string) (err error) {
	return uc.Repo.Delete(id)
}
