package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ViPDanger/dajs/go-api/internal/interfaces/dto"
	"github.com/ViPDanger/dajs/go-api/internal/interfaces/mapper"
	"github.com/ViPDanger/dajs/go-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type playerCharHandler struct {
	Usecase usecase.PlayerCharUsecase
}

func NewPlayerCharHandler(uc usecase.PlayerCharUsecase) *playerCharHandler {
	return &playerCharHandler{Usecase: uc}
}

func (h *playerCharHandler) Get(c *gin.Context) {
	// проверка id header
	clientId, ok := c.Get("client_id")
	if !ok {
		err := errors.New("playerCharHandler.Get(): client_id не найден")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// обращение к Usecase
	objects, err := h.Usecase.Get(c.Request.Context(), string(clientId.(string)))
	if err != nil {
		err = fmt.Errorf("playerCharHandler.Get()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dtos := make([]dto.PlayerCharDTO, len(objects))

	for i := range objects {
		dtos[i] = mapper.ToPlayerCharDTO((*objects[i]))
	}
	// ВЫВОД
	c.JSON(http.StatusOK, dtos)
}

// POST object
func (h *playerCharHandler) New(c *gin.Context) {
	var DTO dto.PlayerCharDTO
	// проверка body запроса
	if err := c.ShouldBindJSON(&DTO); err != nil {
		err = fmt.Errorf("playerCharHandler.New()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "UnmarshalJSON error"})
		return
	}
	object := mapper.ToPlayerCharEntity(DTO)
	clientID, _ := c.Get("client_id")
	object.CreatorID = clientID.(string)
	// Обращение к Usecase
	id, err := h.Usecase.New(c.Request.Context(), &object)
	if err != nil || id == nil {
		err = fmt.Errorf("playerCharHandler.New()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// ВЫВОД
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// DELETE object
func (h *playerCharHandler) Delete(c *gin.Context) {
	id := c.GetHeader("id")
	if id == "" {
		err := errors.New("playerCharHandler.Delete(): No ID header in request")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := h.Usecase.Delete(c.Request.Context(), id)
	if err != nil {
		err = fmt.Errorf("playerCharHandler.Delete()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "DELETE: SUCCESS"})
}

// PUT object
func (h *playerCharHandler) Set(c *gin.Context) {
	var DTO dto.PlayerCharDTO
	if err := c.ShouldBindJSON(&DTO); err != nil {
		err = fmt.Errorf("playerCharHandler.Set()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный JSON"})
		return
	}
	object := mapper.ToPlayerCharEntity(DTO)
	clientID, _ := c.Get("client_id")
	object.CreatorID = clientID.(string)
	err := h.Usecase.Set(c.Request.Context(), &object)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "SET: SUCCESS",
	})
}
