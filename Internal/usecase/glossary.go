package usecase

import (
	"DAJ/Internal/domain/entity"
	"errors"
)

var glossarys = map[string]entity.Glossary{}

type GlossaryUseCase struct {
}

func (charUC *GlossaryUseCase) New(Glossary *entity.Glossary) error {
	if _, exists := glossarys[Glossary.ID]; exists {
		return errors.New("Персонаж с таким ID уже существует")
	}
	glossarys[Glossary.ID] = *Glossary
	return nil
}
func (charUC *GlossaryUseCase) Set(Glossary *entity.Glossary) error {
	if _, exists := glossarys[Glossary.ID]; !exists {
		return errors.New("Персонаж с таким ID не существует")
	}
	glossarys[Glossary.ID] = *Glossary
	return nil
}
func (charUC *GlossaryUseCase) GetByID(id string) (Glossary *entity.Glossary, err error) {
	var exists bool
	var r entity.Glossary
	if r, exists = glossarys[id]; !exists {
		err = errors.New("Персонаж с таким ID не существует")
	}
	Glossary = &r
	return
}
func (charUC *GlossaryUseCase) GetAll() (m *[]entity.Glossary, err error) {
	r := make([]entity.Glossary, len(glossarys))
	var i int
	for k := range glossarys {
		r[i] = glossarys[k]
		i++
	}
	m = &r
	return
}
func (charUC *GlossaryUseCase) Delete(id string) (err error) {
	if _, exists := glossarys[id]; !exists {
		err = errors.New("Персонаж с таким ID не существует")
	}
	delete(glossarys, id)
	return nil
}
