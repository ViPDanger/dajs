package usecase

import "DAJ/Internal/entity"

type CharacterUseCase struct {
}

func New() CharacterUseCase {
	return CharacterUseCase{}
}

func (charUC *CharacterUseCase) New(character entity.Character) error {
	return nil
}
func (charUC *CharacterUseCase) Set(character entity.Character) error {
	return nil
}
func (charUC *CharacterUseCase) Get(id string) (entity.Character, error) {
	return entity.Character{
		ID:   "01",
		Name: "egor",
	}, nil
}
func (charUC *CharacterUseCase) All() ([]entity.Character, error) {
	return []entity.Character{}, nil
}
func (charUC *CharacterUseCase) Delete(id string) error {
	return nil
}
