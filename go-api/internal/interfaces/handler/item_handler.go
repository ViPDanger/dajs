package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/interfaces/dto"
	"github.com/ViPDanger/dajs/go-api/internal/interfaces/mapper"
	"github.com/ViPDanger/dajs/go-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type itemHandler struct {
	UC usecase.ItemUsecase
}

func NewItemHandler(UC usecase.ItemUsecase) *itemHandler {
	return &itemHandler{}
}

// GET object
func (h *itemHandler) Get(c *gin.Context) {
	// проверка id header
	id := c.GetHeader("id")
	if id == "" {
		err := errors.New("itemHandler.Get(): Нет ID в запросе")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// обращение к Usecase
	object, err := h.UC.GetByID(c.Request.Context(), entity.ID(id))
	if err != nil {
		err = fmt.Errorf("itemHandler.Get()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// ВЫВОД
	c.JSON(http.StatusOK, mapper.ToItemDTO(object))
}

func (h *itemHandler) GetByCreatorID(c *gin.Context) {
	// проверка id header
	id := c.GetHeader("creator_id")
	if id == "" {
		err := errors.New("itemHandler.Get(): Нет ID в запросе")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
}

// POST object
func (h *itemHandler) New(c *gin.Context) {
	var DTO dto.ItemDTO
	// проверка body запроса
	if err := c.ShouldBindJSON(&DTO); err != nil {
		err = fmt.Errorf("itemHandler.New()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "UnmarshalJSON error"})
		return
	}
	object := mapper.ToItemEntity(DTO)
	// Обращение к Usecase
	id, err := h.UC.New(c.Request.Context(), object)
	if err != nil || id != nil {
		err = fmt.Errorf("itemHandler.New()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// ВЫВОД
	c.JSON(http.StatusCreated, gin.H{
		"id": id.String(),
	})
}

// GET all objects
func (h *itemHandler) GetAll(c *gin.Context) {
	Objects, err := h.UC.GetAll(c.Request.Context())
	if err != nil {
		err = fmt.Errorf("itemHandler.GetAll()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ObjectsDTO := make([]dto.ItemDTO, len(Objects))
	for i := range Objects {
		ObjectsDTO[i] = mapper.ToItemDTO(Objects[i])
	}
	c.JSON(http.StatusOK, ObjectsDTO)
}

// DELETE object
func (h *itemHandler) Delete(c *gin.Context) {
	id := c.GetHeader("id")
	if id == "" {
		err := errors.New("itemHandler.Delete(): No ID header in request")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := h.UC.Delete(c.Request.Context(), entity.ID(id))
	if err != nil {
		err = fmt.Errorf("itemHandler.Delete()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "DELETE: SUCCESS"})
}

// PUT object
func (h *itemHandler) Set(c *gin.Context) {
	var DTO dto.ItemDTO
	if err := c.ShouldBindJSON(&DTO); err != nil {
		err = fmt.Errorf("itemHandler.Set()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный JSON"})
		return
	}
	object := mapper.ToItemEntity(DTO)
	err := h.UC.Set(c.Request.Context(), object)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "SET: SUCCESS",
	})
}
