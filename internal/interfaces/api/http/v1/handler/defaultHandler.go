package handler

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/interfaces/api/dto"
	"DAJ/internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultHandler[T entity.Identifiable, Tdto any] struct {
	UC       usecase.UseCase[T]
	ToEntity func(Tdto) T
	ToDTO    func(T) Tdto
}

func NewDefaultHandler[T entity.Identifiable, Tdto any](UC usecase.UseCase[T], ToEntity func(Tdto) T, ToDTO func(T) Tdto) *DefaultHandler[T, Tdto] {
	return &DefaultHandler[T, Tdto]{UC: UC, ToEntity: ToEntity, ToDTO: ToDTO}
}

// GET character
func (h *DefaultHandler[T, Tdto]) Get(c *gin.Context) {
	id := c.GetHeader("id")
	if id == "" {
		err := errors.New("Нет ID в запросе")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	object, err := h.UC.GetByID(id)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, h.ToDTO(*object))
}

// POST character
func (h *DefaultHandler[T, Tdto]) New(c *gin.Context) {
	var DTO Tdto
	if err := c.ShouldBindJSON(&DTO); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, dto.Error{Error: "Некорректный JSON"})
		return
	}
	object := h.ToEntity(DTO)
	err := h.UC.New(&object)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.Message{
		Message: "Character was created",
	})
}

// GET all chatacter
func (h *DefaultHandler[T, Tdto]) GetAll(c *gin.Context) {
	Objects, err := h.UC.GetAll()
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
		return
	}
	ObjectsDTO := make([]Tdto, len(Objects))
	for i := range Objects {
		ObjectsDTO[i] = h.ToDTO(Objects[i])
	}
	c.JSON(http.StatusOK, ObjectsDTO)
}

// DELETE character
func (h *DefaultHandler[T, Tdto]) Delete(c *gin.Context) {
	id := c.GetHeader("id")
	if id == "" {
		err := errors.New("Нет ID в запросе")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := h.UC.Delete(id)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, dto.Message{Message: "Character with id " + id + "deleted"})
}

// PUT character
func (h *DefaultHandler[T, Tdto]) Set(c *gin.Context) {
	var DTO Tdto
	if err := c.ShouldBindJSON(&DTO); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, dto.Error{Error: "Некорректный JSON"})
		return
	}
	object := h.ToEntity(DTO)
	err := h.UC.Set(&object)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.Message{
		Message: "Успешно",
	})
}
