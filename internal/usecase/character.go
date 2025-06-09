package usecase

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
)

type CharacterUseCase struct {
	CharRepository repository.Repository[entity.Character]
}

func (uc *CharacterUseCase) New(character *entity.Character) error {
	return uc.CharRepository.New(character.ID, character)
}
func (uc *CharacterUseCase) Set(character *entity.Character) error {
	return uc.CharRepository.Set(character.ID, character)
}
func (uc *CharacterUseCase) GetByID(id string) (r *entity.Character, err error) {
	return uc.CharRepository.GetByID(id)
}
func (uc *CharacterUseCase) GetAll() (r *[]entity.Character, err error) {
	return uc.CharRepository.GetAll()
}
func (uc *CharacterUseCase) Delete(id string) (err error) {
	return uc.CharRepository.Delete(id)
}
