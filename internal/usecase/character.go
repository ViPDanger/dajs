package usecase

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
)

type CharacterUseCase struct {
	Repo repository.Repository[entity.Character]
}

func (uc *CharacterUseCase) New(character *entity.Character) error {
	err := uc.Repo.Insert(character)
	return err
}
func (uc *CharacterUseCase) Set(character *entity.Character) error {
	return uc.Repo.Update(character)
}
func (uc *CharacterUseCase) GetByID(id string) (r *entity.Character, err error) {
	return uc.Repo.GetByID(id)
}
func (uc *CharacterUseCase) GetAll() (r []entity.Character, err error) {
	return uc.Repo.GetAll()
}
func (uc *CharacterUseCase) Delete(id string) (err error) {
	return uc.Repo.Delete(id)
}
