package json

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
	"DAJ/internal/infastructure/json/jsonDTO"
	"DAJ/internal/infastructure/json/jsonMapper"
)

type characterRepository struct {
	defaultJSONRepository[entity.Character, jsonDTO.CharacterDTO]
}

func NewCharacterRepository(filepath string) (repository.Repository[entity.Character], error) {
	r := characterRepository{}
	defaultRepository, err := NewJSONRepository(filepath, jsonMapper.ToCharacterDTO, jsonMapper.ToCharacterEntity, r.characterPathFunc)
	if err != nil {
		return nil, err
	}
	defaultRepository.pathFunc = func(character *entity.Character) string {
		return defaultRepository.fileDirectory + character.Name + defaultFileType
	}
	r.defaultJSONRepository = *defaultRepository
	return &r, nil
}
func (r *characterRepository) characterPathFunc(c *entity.Character) string {
	return r.fileDirectory + c.Name + defaultFileType
}
