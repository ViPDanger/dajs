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

type playerHandler struct {
	UC usecase.PlayerUsecase
}

func NewPlayerCharacterHandler(UC usecase.PlayerUsecase) *playerHandler {
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

func (h *playerHandler) GetByCreatorID(c *gin.Context) {
	// проверка id header
	clientId, ok := c.Get("client_id")
	if !ok {
		err := errors.New("playerHandler.Get(): Нет ID в запросе")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// обращение к Usecase
	objects, err := h.UC.GetByCreatorID(c.Request.Context(), entity.ID(clientId.(string)))
	if err != nil {
		err = fmt.Errorf("playerHandler.Get()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dtos := make([]dto.PlayerCharacterDTO, len(objects))
	for i := range objects {
		dtos[i] = mapper.ToPlayerDTO(*(objects[i]))
	}
	// ВЫВОД
	c.JSON(http.StatusOK, dtos)
}

// POST object
func (h *playerHandler) New(c *gin.Context) {
	var DTO dto.PlayerCharacterDTO
	// проверка body запроса
	if err := c.ShouldBindJSON(&DTO); err != nil {
		err = fmt.Errorf("playerHandler.New()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "UnmarshalJSON error"})
		return
	}
	object := mapper.ToPlayerEntity(DTO)
	clientID, _ := c.Get("client_id")
	object.CreatorID = clientID.(entity.ID).String()
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
	ObjectsDTO := make([]dto.PlayerCharacterDTO, len(Objects))
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
	c.JSON(http.StatusCreated, gin.H{"message": "DELETE: SUCCESS"})
}

// PUT object
func (h *playerHandler) Set(c *gin.Context) {
	var DTO dto.PlayerCharacterDTO
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
