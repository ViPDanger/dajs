package json

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
	"DAJ/internal/infastructure/json/jsonDTO"
	"DAJ/internal/infastructure/json/jsonMapper"
)

type abilityJSONRepository struct {
	defaultJSONRepository[entity.Ability, jsonDTO.AbilityDTO]
}

func NewAbilityRepository(filepath string) (repository.Repository[entity.Ability], error) {
	r := abilityJSONRepository{}
	defaultRepository, err := NewJSONRepository(filepath, jsonMapper.ToAbilityDTO, jsonMapper.ToAbilityEntity, r.AbilityPathFunc)
	if err != nil {
		return nil, err
	}
	r.defaultJSONRepository = *defaultRepository
	return &r, nil
}
func (r *abilityJSONRepository) AbilityPathFunc(u *entity.Ability) string {
	return r.fileDirectory + "/" + u.GetID() + defaultFileType
}
