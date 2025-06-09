package usecase

import (
	"DAJ/Internal/domain/entity"
	"DAJ/Internal/domain/repository"
	"errors"
)

type UseCase[T any] interface {
	New(item *T) error
	GetByID(id string) (*T, error)
	GetAll() (*[]T, error)
	Set(item *T) error
	Delete(id string) error
}

var characters = map[string]entity.Character{}

type CharacterUseCase struct {
	charRepository repository.Repository[entity.Character]
}

func (charUC *CharacterUseCase) New(character *entity.Character) error {
	if _, exists := characters[character.ID]; exists {
		return errors.New("Персонаж с таким ID уже существует")
	}
	characters[character.ID] = *character
	return nil
}
func (charUC *CharacterUseCase) Set(character *entity.Character) error {
	if _, exists := characters[character.ID]; !exists {
		return errors.New("Персонаж с таким ID не существует")
	}
	characters[character.ID] = *character
	return nil
}
func (charUC *CharacterUseCase) GetByID(id string) (r *entity.Character, err error) {
	if character, exists := characters[id]; !exists {
		err = errors.New("Персонаж с таким ID не существует")
	} else {
		r = &character
	}
	return
}
func (charUC *CharacterUseCase) GetAll() (r *[]entity.Character, err error) {
	m := make([]entity.Character, len(characters))
	var i int
	for _, character := range characters {
		m[i] = character
	}
	r = &m
	return
}
func (charUC *CharacterUseCase) Delete(id string) (err error) {
	if _, exists := characters[id]; !exists {
		err = errors.New("Персонаж с таким ID не существует")
	}
	delete(characters, id)
	return nil
}
