package json

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
	"DAJ/internal/infastructure/json/jsonDTO"
	"DAJ/internal/infastructure/json/jsonMapper"
)

type statusJSONRepository struct {
	defaultJSONRepository[entity.Status, jsonDTO.StatusDTO]
}

func NewStatusRepository(filepath string) (repository.Repository[entity.Status], error) {
	r := statusJSONRepository{}
	defaultRepository, err := NewJSONRepository(filepath, jsonMapper.ToStatusDTO, jsonMapper.ToStatusEntity, r.StatusPathFunc)
	if err != nil {
		return nil, err
	}
	r.defaultJSONRepository = *defaultRepository
	return &r, nil
}
func (r *statusJSONRepository) StatusPathFunc(u *entity.Status) string {
	return r.fileDirectory + u.GetID() + defaultFileType
}
