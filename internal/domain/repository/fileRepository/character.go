package fileRepository

import (
	"DAJ/internal/domain/entity"
)

type CharacterFileRepository struct {
	fileRepository[entity.Character]
}

func NewCharacterFileRepository(filepath string) (*CharacterFileRepository, error) {
	repository, err := NewFileRepository[entity.Character](filepath)
	if err != nil {
		return nil, err
	}
	return &CharacterFileRepository{fileRepository: *repository}, nil
}

func (r *CharacterFileRepository) Insert(item *entity.Character) error {
	id := item.ID
	r.ItemIds[id] = item.Name
	return r.file.Add(r.fileDirectory+r.ItemIds[id]+".json", item)
}

func (r *CharacterFileRepository) Update(item *entity.Character) error {
	id := item.ID
	err := r.file.Patch(r.fileDirectory+r.ItemIds[id]+".json", item)
	//r.ItemIds[id] = item.Name
	return err
}
