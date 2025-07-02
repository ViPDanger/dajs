package usecase

import (
	"fmt"

	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/domain/repository"
)

type characterUsecase struct {
	UseCase[entity.Character]
	itemUseCase UseCase[entity.Item]
}

func NewCharacterUseCase(Repository repository.Repository[entity.Character], itemRepository UseCase[entity.Item]) UseCase[entity.Character] {
	return &characterUsecase{UseCase: NewDefaultUsecase(Repository), itemUseCase: itemRepository}
}

func (u characterUsecase) getInventory(Inventorys []entity.CharacterInventory) error {
	for _, inventory := range Inventorys {
		s := make([]string, 0)

		for i := range inventory.Items {
			if inventory.Items[i].ID != nil {
				s = append(s, *inventory.Items[i].ID)
			}

		}
		items, err := u.itemUseCase.GetArray(s)
		if err != nil {
			return fmt.Errorf("NewCharacterUsecase()/%w", err)
		}
		for i, j := 0, 0; j < len(items); i++ {
			if inventory.Items[i].ID != nil {
				inventory.Items[i].Item = items[j]
				j++
			}

		}
	}

	return nil
}
func (u characterUsecase) GetAll() ([]entity.Character, error) {
	objects, err := u.UseCase.GetAll()
	if err != nil {
		return nil, fmt.Errorf("characterUseCase.GetAll()/%w", err)
	}

	for i := range objects {
		if err := u.getInventory(objects[i].Inventory); err != nil {
			return nil, fmt.Errorf("characterUseCase.GetAll()/%w", err)
		}
		fmt.Println(objects[i].Inventory)
	}

	return objects, nil
}
func (u characterUsecase) GetArray(ids []string) ([]entity.Character, error) {
	objects, err := u.UseCase.GetArray(ids)
	if err != nil {
		return nil, fmt.Errorf("characterUseCase.GetArray()/%w", err)
	}
	for i := range objects {
		if err := u.getInventory(objects[i].Inventory); err != nil {
			return nil, fmt.Errorf("characterUseCase.GetArray()/%w", err)
		}
	}
	return objects, nil
}

func (u characterUsecase) GetByID(id string) (*entity.Character, error) {
	object, err := u.UseCase.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("characterUseCase.GetById()/%w", err)
	}
	if err := u.getInventory(object.Inventory); err != nil {
		return nil, fmt.Errorf("characterUseCase.GetById()/%w", err)
	}
	return object, nil
}
func (u characterUsecase) New(object *entity.Character) error {
	if err := u.getInventory(object.Inventory); err != nil {
		return fmt.Errorf("characterUseCase.New()/%w", err)
	}
	if err := u.UseCase.New(object); err != nil {
		return fmt.Errorf("characterUseCase.New()/%w", err)
	}
	return nil
}
func (u characterUsecase) Set(object *entity.Character) error {
	if err := u.getInventory(object.Inventory); err != nil {
		return fmt.Errorf("characterUseCase.Set()/%w", err)
	}
	if err := u.UseCase.Set(object); err != nil {
		return fmt.Errorf("characterUseCase.Set()/%w", err)
	}
	return nil
}
