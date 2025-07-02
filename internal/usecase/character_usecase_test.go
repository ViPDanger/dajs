package usecase_test

import (
	"errors"
	"testing"

	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// === Моки ===

type MockUseCase[T entity.Identifiable] struct {
	mock.Mock
}

func (m *MockUseCase[T]) New(item *T) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockUseCase[T]) GetByID(id string) (*T, error) {
	args := m.Called(id)
	return args.Get(0).(*T), args.Error(1)
}

func (m *MockUseCase[T]) GetArray(ids []string) ([]T, error) {
	args := m.Called(ids)
	return args.Get(0).([]T), args.Error(1)
}

func (m *MockUseCase[T]) GetAll() ([]T, error) {
	args := m.Called()
	return args.Get(0).([]T), args.Error(1)
}

func (m *MockUseCase[T]) Set(item *T) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockUseCase[T]) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// === Тесты ===

func TestCharacterUsecase_GetByID(t *testing.T) {
	mockCharRepo := new(MockRepository[entity.Character])
	mockItemUC := new(MockUseCase[entity.Item])
	charID := "char1"
	itemID := "item1"
	var item entity.Item
	item = &entity.SimpleItem{ID: itemID}
	char := entity.Character{
		ID: charID,
		Inventory: []entity.CharacterInventory{
			{
				Name: "Сумка",
				Items: []entity.CharacterItem{
					{ID: &itemID},
				},
			},
		},
	}

	mockCharRepo.On("GetByID", charID).Return(&char, nil)
	mockItemUC.On("GetArray", []string{itemID}).Return([]entity.Item{item}, nil)

	characterUC := usecase.NewCharacterUseCase(mockCharRepo, mockItemUC)
	result, err := characterUC.GetByID(charID)
	assert.NoError(t, err)
	assert.Equal(t, charID, result.ID)
	assert.Equal(t, itemID, result.Inventory[0].Items[0].Item.GetID())

	mockCharRepo.AssertExpectations(t)
	mockItemUC.AssertExpectations(t)
}

func TestCharacterUsecase_GetByID_ErrorFromItemUsecase(t *testing.T) {
	mockCharRepo := new(MockRepository[entity.Character])
	mockItemUC := new(MockUseCase[entity.Item])

	charID := "char2"
	itemID := "item2"
	var item entity.Item
	item = &entity.SimpleItem{ID: itemID}
	char := entity.Character{
		ID: charID,
		Inventory: []entity.CharacterInventory{
			{
				Name: "Инвентарь",
				Items: []entity.CharacterItem{
					{ID: &itemID},
				},
			},
		},
	}

	mockCharRepo.On("GetByID", charID).Return(&char, nil)
	mockItemUC.On("GetArray", []string{itemID}).Return([]entity.Item{item}, errors.New("item fetch error"))

	characterUC := usecase.NewCharacterUseCase(mockCharRepo, mockItemUC)

	result, err := characterUC.GetByID(charID)
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "item fetch error")
}

func TestCharacterUsecase_GetAll(t *testing.T) {
	mockCharUC := new(MockRepository[entity.Character])
	mockItemUC := new(MockUseCase[entity.Item])

	itemID := "item1"
	item := entity.SimpleItem{ID: itemID}
	char := entity.Character{
		ID: "char1",
		Inventory: []entity.CharacterInventory{
			{
				Name: "Сумка",
				Items: []entity.CharacterItem{
					{ID: &itemID},
				},
			},
		},
	}

	mockCharUC.On("GetAll").Return([]entity.Character{char}, nil)
	mockItemUC.On("GetArray", []string{itemID}).Return([]entity.Item{item}, nil)

	characterUC := usecase.NewCharacterUseCase(mockCharUC, mockItemUC)

	result, err := characterUC.GetAll()
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, itemID, result[0].Inventory[0].Items[0].Item.GetID())
}

func TestCharacterUsecase_GetArray(t *testing.T) {
	mockCharRepo := new(MockRepository[entity.Character])
	mockItemUC := new(MockUseCase[entity.Item])

	itemID := "itemA"
	item := entity.SimpleItem{ID: itemID}
	char := entity.Character{
		ID: "charA",
		Inventory: []entity.CharacterInventory{
			{
				Name: "Сундук",
				Items: []entity.CharacterItem{
					{ID: &itemID},
				},
			},
		},
	}

	mockCharRepo.On("GetArray", []string{"charA"}).Return([]entity.Character{char}, nil)
	mockItemUC.On("GetArray", []string{itemID}).Return([]entity.Item{item}, nil)

	characterUC := usecase.NewCharacterUseCase(mockCharRepo, mockItemUC)

	result, err := characterUC.GetArray([]string{"charA"})
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, itemID, result[0].Inventory[0].Items[0].Item.GetID())
}

func TestCharacterUsecase_New(t *testing.T) {
	mockCharRepo := new(MockRepository[entity.Character])
	mockItemUC := new(MockUseCase[entity.Item])

	itemID := "itemN"
	item := entity.SimpleItem{ID: itemID}
	char := entity.Character{
		ID: "charN",
		Inventory: []entity.CharacterInventory{
			{
				Name: "Карман",
				Items: []entity.CharacterItem{
					{ID: &itemID},
				},
			},
		},
	}

	mockItemUC.On("GetArray", []string{itemID}).Return([]entity.Item{&item}, nil)
	mockItemUC.On("GetArray", []string{itemID}).Return([]entity.Item{&item}, nil)
	mockCharRepo.On("New", &char).Return(nil)

	characterUC := usecase.NewCharacterUseCase(mockCharRepo, mockItemUC)

	err := characterUC.New(&char)
	assert.NoError(t, err)
	assert.Equal(t, itemID, char.Inventory[0].Items[0].Item.GetID())
}

func TestCharacterUsecase_Set(t *testing.T) {
	mockCharUC := new(MockRepository[entity.Character])
	mockItemUC := new(MockUseCase[entity.Item])

	itemID := "itemS"
	item := entity.SimpleItem{ID: itemID}
	char := entity.Character{
		ID: "charS",
		Inventory: []entity.CharacterInventory{
			{
				Name: "Склад",
				Items: []entity.CharacterItem{
					{ID: &itemID},
				},
			},
		},
	}

	mockItemUC.On("GetArray", []string{itemID}).Return([]entity.Item{item}, nil)
	mockCharUC.On("Set", &char).Return(nil)

	characterUC := usecase.NewCharacterUseCase(mockCharUC, mockItemUC)

	err := characterUC.Set(&char)
	assert.NoError(t, err)
	assert.Equal(t, itemID, char.Inventory[0].Items[0].Item.GetID())
}
