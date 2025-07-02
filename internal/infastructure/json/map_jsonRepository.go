package json

import (
	"fmt"

	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/domain/repository"
	"github.com/ViPDanger/dajs/internal/infastructure/json/jsonDTO"
	"github.com/ViPDanger/dajs/internal/infastructure/json/jsonMapper"
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
