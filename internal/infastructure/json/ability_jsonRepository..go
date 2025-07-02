package json

import (
	"fmt"

	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/domain/repository"
	"github.com/ViPDanger/dajs/internal/infastructure/json/jsonDTO"
	"github.com/ViPDanger/dajs/internal/infastructure/json/jsonMapper"
)

type abilityJSONRepository struct {
	defaultJSONRepository[entity.Ability, jsonDTO.AbilityDTO]
}

func NewAbilityRepository(filepath string) (repository.Repository[entity.Ability], error) {
	r := abilityJSONRepository{}
	defaultRepository, err := NewJSONRepository(filepath, jsonMapper.ToAbilityDTO, jsonMapper.ToAbilityEntity, r.AbilityPathFunc)
	if err != nil {
		return nil, fmt.Errorf("NewAbilityRepository()/%w", err)
	}
	r.defaultJSONRepository = *defaultRepository
	return &r, nil
}
func (r *abilityJSONRepository) AbilityPathFunc(u *entity.Ability) string {
	return r.fileDirectory + "/" + u.GetID() + defaultFileType
}
