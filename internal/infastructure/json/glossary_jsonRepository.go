package json

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
	"DAJ/internal/infastructure/json/jsonDTO"
	"DAJ/internal/infastructure/json/jsonMapper"
)

type glossaryRepository struct {
	defaultJSONRepository[entity.Glossary, jsonDTO.GlossaryDTO]
}

func NewGlossaryRepository(filepath string) (repository.Repository[entity.Glossary], error) {
	r := glossaryRepository{}
	defaultRepository, err := NewJSONRepository(filepath, jsonMapper.ToGlossaryDTO, jsonMapper.ToGlossaryEntity, r.GlossaryPathFunc)
	if err != nil {
		return nil, err
	}
	defaultRepository.pathFunc = func(Glossary *entity.Glossary) string {
		return defaultRepository.fileDirectory + Glossary.ID + defaultFileType
	}
	r.defaultJSONRepository = *defaultRepository
	return &r, nil
}
func (r *glossaryRepository) GlossaryPathFunc(u *entity.Glossary) string {
	return r.fileDirectory + u.ID + defaultFileType
}
