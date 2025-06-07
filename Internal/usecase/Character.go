package usecase

import (
	"DAJ/Internal/domain/entity"
	"errors"
)

var characters = map[string]entity.Character{}

type CharacterUseCase struct {
}

func (charUC *CharacterUseCase) New(character entity.Character) error {
	if _, exists := characters[character.ID]; exists {
		return errors.New("Персонаж с таким ID уже существует")
	}
	characters[character.ID] = character
	return nil
}
func (charUC *CharacterUseCase) Set(character entity.Character) error {
	if _, exists := characters[character.ID]; !exists {
		return errors.New("Персонаж с таким ID не существует")
	}
	characters[character.ID] = character
	return nil
}
func (charUC *CharacterUseCase) Get(id string) (character entity.Character, err error) {
	var exists bool
	if character, exists = characters[id]; !exists {
		err = errors.New("Персонаж с таким ID не существует")
	}
	return
}
func (charUC *CharacterUseCase) GetAll() (m []entity.Character, err error) {
	for i := range characters {
		m = append(m, characters[i])
	}
	return
}
func (charUC *CharacterUseCase) Delete(id string) (err error) {
	if _, exists := characters[id]; !exists {
		err = errors.New("Персонаж с таким ID не существует")
	}
	delete(characters, id)
	return nil
}
