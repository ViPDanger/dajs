package usecase_test

import (
	"errors"
	"testing"

	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ==== Entity =====

type TestItem struct {
	ID string
}

func (t TestItem) GetID() string {
	return t.ID
}

// ==== Mock Repository ====

type MockRepository[T entity.Identifiable] struct {
	mock.Mock
}

func (m *MockRepository[T]) Insert(item *T) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockRepository[T]) Update(item *T) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockRepository[T]) GetByID(id string) (*T, error) {
	args := m.Called(id)
	return args.Get(0).(*T), args.Error(1)
}

func (m *MockRepository[T]) GetArray(ids []string) ([]T, error) {
	args := m.Called(ids)
	return args.Get(0).([]T), args.Error(1)
}

func (m *MockRepository[T]) GetAll() ([]T, error) {
	args := m.Called()
	return args.Get(0).([]T), args.Error(1)
}

func (m *MockRepository[T]) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// ==== Tests ====

func TestDefaultUseCase_New(t *testing.T) {
	mockRepo := new(MockRepository[TestItem])
	uc := usecase.NewDefaultUsecase(mockRepo)

	item := TestItem{ID: "1"}
	mockRepo.On("Insert", &item).Return(nil)

	err := uc.New(&item)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDefaultUseCase_New_Nil(t *testing.T) {
	mockRepo := new(MockRepository[TestItem])
	uc := usecase.NewDefaultUsecase(mockRepo)

	err := uc.New(nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Nill pointer")
}

func TestDefaultUseCase_New_ErrorFromRepo(t *testing.T) {
	mockRepo := new(MockRepository[TestItem])
	uc := usecase.NewDefaultUsecase(mockRepo)

	item := TestItem{ID: "1"}
	mockRepo.On("Insert", &item).Return(errors.New("insert failed"))

	err := uc.New(&item)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "insert failed")
}

// -- Set --

func TestDefaultUseCase_Set(t *testing.T) {
	mockRepo := new(MockRepository[TestItem])
	uc := usecase.NewDefaultUsecase(mockRepo)

	item := TestItem{ID: "2"}
	mockRepo.On("Update", &item).Return(nil)

	err := uc.Set(&item)
	assert.NoError(t, err)
}

func TestDefaultUseCase_Set_Nil(t *testing.T) {
	mockRepo := new(MockRepository[TestItem])
	uc := usecase.NewDefaultUsecase[TestItem](mockRepo)

	err := uc.Set(nil)
	assert.Error(t, err)
}

func TestDefaultUseCase_Set_Error(t *testing.T) {
	mockRepo := new(MockRepository[TestItem])
	uc := usecase.NewDefaultUsecase[TestItem](mockRepo)

	item := TestItem{ID: "2"}
	mockRepo.On("Update", &item).Return(errors.New("update error"))

	err := uc.Set(&item)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "update error")
}

// -- GetByID --

func TestDefaultUseCase_GetByID(t *testing.T) {
	mockRepo := new(MockRepository[TestItem])
	uc := usecase.NewDefaultUsecase[TestItem](mockRepo)

	item := TestItem{ID: "3"}
	mockRepo.On("GetByID", "3").Return(&item, nil)

	result, err := uc.GetByID("3")
	assert.NoError(t, err)
	assert.Equal(t, "3", result.ID)
}

func TestDefaultUseCase_GetByID_Empty(t *testing.T) {
	mockRepo := new(MockRepository[TestItem])
	uc := usecase.NewDefaultUsecase[TestItem](mockRepo)

	result, err := uc.GetByID("")
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestDefaultUseCase_GetByID_Error(t *testing.T) {
	mockRepo := new(MockRepository[TestItem])
	uc := usecase.NewDefaultUsecase[TestItem](mockRepo)

	mockRepo.On("GetByID", "bad-id").Return((*TestItem)(nil), errors.New("not found"))

	result, err := uc.GetByID("bad-id")
	assert.Error(t, err)
	assert.Nil(t, result)
}

// -- GetArray --

func TestDefaultUseCase_GetArray(t *testing.T) {
	mockRepo := new(MockRepository[TestItem])
	uc := usecase.NewDefaultUsecase[TestItem](mockRepo)

	ids := []string{"1", "2"}
	items := []TestItem{{ID: "1"}, {ID: "2"}}
	mockRepo.On("GetArray", ids).Return(items, nil)

	result, err := uc.GetArray(ids)
	assert.NoError(t, err)
	assert.Len(t, result, 2)
}

func TestDefaultUseCase_GetArray_Error(t *testing.T) {
	mockRepo := new(MockRepository[TestItem])
	uc := usecase.NewDefaultUsecase[TestItem](mockRepo)

	ids := []string{"1"}
	mockRepo.On("GetArray", ids).Return(nil, errors.New("array error"))

	result, err := uc.GetArray(ids)
	assert.Error(t, err)
	assert.Nil(t, result)
}

// -- GetAll --

func TestDefaultUseCase_GetAll(t *testing.T) {
	mockRepo := new(MockRepository[TestItem])
	uc := usecase.NewDefaultUsecase[TestItem](mockRepo)

	expected := []TestItem{{ID: "1"}, {ID: "2"}}
	mockRepo.On("GetAll").Return(expected, nil)

	result, err := uc.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestDefaultUseCase_GetAll_Error(t *testing.T) {
	mockRepo := new(MockRepository[TestItem])
	uc := usecase.NewDefaultUsecase[TestItem](mockRepo)

	mockRepo.On("GetAll").Return(nil, errors.New("getall error"))

	result, err := uc.GetAll()
	assert.Error(t, err)
	assert.Nil(t, result)
}

// -- Delete --

func TestDefaultUseCase_Delete(t *testing.T) {
	mockRepo := new(MockRepository[TestItem])
	uc := usecase.NewDefaultUsecase[TestItem](mockRepo)

	mockRepo.On("Delete", "1").Return(nil)

	err := uc.Delete("1")
	assert.NoError(t, err)
}

func TestDefaultUseCase_Delete_Error(t *testing.T) {
	mockRepo := new(MockRepository[TestItem])
	uc := usecase.NewDefaultUsecase[TestItem](mockRepo)

	mockRepo.On("Delete", "bad").Return(errors.New("delete error"))

	err := uc.Delete("bad")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "delete error")
}
