package json

import (
	"fmt"

	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/domain/repository"
	"github.com/ViPDanger/dajs/internal/infastructure/json/jsonDTO"
	"github.com/ViPDanger/dajs/internal/infastructure/json/jsonMapper"
)

type glossaryJSONRepository struct {
	defaultJSONRepository[entity.Glossary, jsonDTO.GlossaryDTO]
}

func NewGlossaryRepository(filepath string) (repository.Repository[entity.Glossary], error) {
	r := glossaryJSONRepository{}
	defaultRepository, err := NewJSONRepository(filepath, jsonMapper.ToGlossaryDTO, jsonMapper.ToGlossaryEntity, r.GlossaryPathFunc)
	if err != nil {
		return nil, fmt.Errorf("NewGlossaryRepository()/%w", err)
	}
	defaultRepository.pathFunc = r.GlossaryPathFunc
	r.defaultJSONRepository = *defaultRepository
	return &r, nil
}
func (r *glossaryJSONRepository) GlossaryPathFunc(u *entity.Glossary) string {
	return r.fileDirectory + "/" + u.ID + defaultFileType
}
