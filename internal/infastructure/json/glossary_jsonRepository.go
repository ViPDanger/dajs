package json

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
	"DAJ/internal/infastructure/json/jsonDTO"
	"DAJ/internal/infastructure/json/jsonMapper"
)

type glossaryJSONRepository struct {
	defaultJSONRepository[entity.Glossary, jsonDTO.GlossaryDTO]
}

func NewGlossaryRepository(filepath string) (repository.Repository[entity.Glossary], error) {
	r := glossaryJSONRepository{}
	defaultRepository, err := NewJSONRepository(filepath, jsonMapper.ToGlossaryDTO, jsonMapper.ToGlossaryEntity, r.GlossaryPathFunc)
	if err != nil {
		return nil, err
	}
	defaultRepository.pathFunc = r.GlossaryPathFunc
	r.defaultJSONRepository = *defaultRepository
	return &r, nil
}
func (r *glossaryJSONRepository) GlossaryPathFunc(u *entity.Glossary) string {
	return r.fileDirectory + "/" + u.ID + defaultFileType
}
