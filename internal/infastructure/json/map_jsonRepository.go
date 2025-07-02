package json

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
	"DAJ/internal/infastructure/json/jsonDTO"
	"DAJ/internal/infastructure/json/jsonMapper"
	"fmt"
)

// НАПИСААААТь
type mapJSONRepository struct {
	defaultJSONRepository[entity.Map, jsonDTO.MapDTO]
}

func NewMapRepository(filepath string) (repository.Repository[entity.Map], error) {
	r := mapJSONRepository{}
	defaultRepository, err := NewJSONRepository(filepath, jsonMapper.ToMapDTO, jsonMapper.ToMapEntity, r.mapPathFunc)
	if err != nil {
		return nil, fmt.Errorf("NewMapRepository()/%w", err)
	}
	r.defaultJSONRepository = *defaultRepository
	return &r, nil
}
func (r *mapJSONRepository) mapPathFunc(u *entity.Map) string {
	return r.fileDirectory + u.GetID() + defaultFileType
}
