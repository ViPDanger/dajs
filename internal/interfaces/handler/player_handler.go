package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/interfaces/dto"
	"github.com/ViPDanger/dajs/internal/interfaces/mapper"
	"github.com/ViPDanger/dajs/internal/usecase"
	"github.com/gin-gonic/gin"
)

type playerHandler struct {
	UC usecase.PlayerUsecase
}

func NewPlayerHandler(UC usecase.PlayerUsecase) *playerHandler {
	return &playerHandler{}
}

// GET object
func (h *playerHandler) Get(c *gin.Context) {
	// проверка id header
	id := c.GetHeader("id")
	if id == "" {
		err := errors.New("playerHandler.Get(): Нет ID в запросе")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// обращение к Usecase
	object, err := h.UC.GetByID(c.Request.Context(), entity.ID(id))
	if err != nil {
		err = fmt.Errorf("playerHandler.Get()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// ВЫВОД
	c.JSON(http.StatusOK, mapper.ToPlayerDTO(*object))
}

// POST object
func (h *playerHandler) New(c *gin.Context) {
	var DTO dto.PlayerDTO
	// проверка body запроса
	if err := c.ShouldBindJSON(&DTO); err != nil {
		err = fmt.Errorf("playerHandler.New()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "UnmarshalJSON error"})
		return
	}
	object := mapper.ToPlayerEntity(DTO)
	// Обращение к Usecase
	id, err := h.UC.New(c.Request.Context(), &object)
	if err != nil || id != nil {
		err = fmt.Errorf("playerHandler.New()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// ВЫВОД
	c.JSON(http.StatusOK, gin.H{
		"id": id.String(),
	})
}

// GET all objects
func (h *playerHandler) GetAll(c *gin.Context) {
	Objects, err := h.UC.GetAll(c.Request.Context())
	if err != nil {
		err = fmt.Errorf("playerHandler.GetAll()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ObjectsDTO := make([]dto.PlayerDTO, len(Objects))
	for i := range Objects {
		ObjectsDTO[i] = mapper.ToPlayerDTO(*Objects[i])
	}
	c.JSON(http.StatusOK, ObjectsDTO)
}

// DELETE object
func (h *playerHandler) Delete(c *gin.Context) {
	id := c.GetHeader("id")
	if id == "" {
		err := errors.New("playerHandler.Delete(): No ID header in request")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := h.UC.Delete(c.Request.Context(), entity.ID(id))
	if err != nil {
		err = fmt.Errorf("playerHandler.Delete()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "DELETE: SUCCESS"})
}

// PUT object
func (h *playerHandler) Set(c *gin.Context) {
	var DTO dto.PlayerDTO
	if err := c.ShouldBindJSON(&DTO); err != nil {
		err = fmt.Errorf("playerHandler.Set()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный JSON"})
		return
	}
	object := mapper.ToPlayerEntity(DTO)
	err := h.UC.Set(c.Request.Context(), &object)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "SET: SUCCESS",
	})
}
