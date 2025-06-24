package json

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
	"DAJ/internal/infastructure/json/jsonDTO"
	"DAJ/internal/infastructure/json/jsonMapper"
)

type characterJSONRepository struct {
	defaultJSONRepository[entity.Character, jsonDTO.CharacterDTO]
}

func NewCharacterRepository(filepath string) (repository.Repository[entity.Character], error) {
	r := characterJSONRepository{}
	defaultRepository, err := NewJSONRepository(filepath, jsonMapper.ToCharacterDTO, jsonMapper.ToCharacterEntity, r.characterPathFunc)
	if err != nil {
		return nil, err
	}
	r.defaultJSONRepository = *defaultRepository
	return &r, nil
}
func (r *characterJSONRepository) characterPathFunc(c *entity.Character) string {
	return r.fileDirectory + c.Name + defaultFileType
}
