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

type characterHandler struct {
	Usecase usecase.CharacterUsecase
}

func NewCharacterHandler(uc usecase.CharacterUsecase) *characterHandler {
	return &characterHandler{Usecase: uc}
}

// GET object
func (h *characterHandler) Get(c *gin.Context) {
	// проверка id header
	id := c.GetHeader("id")
	if id == "" {
		err := errors.New("characterHandler.Get():No id header in request")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// обращение к Usecase
	object, err := h.Usecase.GetByID(c.Request.Context(), entity.ID(id))
	if err != nil {
		err = fmt.Errorf("characterHandler.Get()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// ВЫВОД
	c.JSON(http.StatusOK, mapper.ToCharacterDTO(*object))
}

func (h *characterHandler) GetByCreatorID(c *gin.Context) {
	// проверка clientId header
	clientId, ok := c.Get("client_id")
	if !ok {
		err := errors.New("characterHandler.Get(): client_id не найден")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// обращение к Usecase
	objects, err := h.Usecase.GetByCreatorID(c.Request.Context(), clientId.(entity.ID))
	if err != nil {
		err = fmt.Errorf("characterHandler.Get()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dtos := make([]dto.CharacterDTO, len(objects))

	for i := range objects {
		dtos[i] = mapper.ToCharacterDTO(*(objects[i]))
	}
	// ВЫВОД
	c.JSON(http.StatusOK, dtos)
}

// POST object
func (h *characterHandler) New(c *gin.Context) {
	var DTO dto.CharacterDTO
	// проверка body запроса
	if err := c.ShouldBindJSON(&DTO); err != nil {
		err = fmt.Errorf("characterHandler.New()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "UnmarshalJSON error"})
		return
	}
	object := mapper.ToCharacterEntity(DTO)
	clientID, _ := c.Get("client_id")
	object.CreatorID = clientID.(entity.ID).String()
	// Обращение к Usecase
	id, err := h.Usecase.New(c.Request.Context(), &object)
	if err != nil || id == nil {
		err = fmt.Errorf("characterHandler.New()/%w", err)
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
func (h *characterHandler) GetAll(c *gin.Context) {
	Objects, err := h.Usecase.GetAll(c.Request.Context())
	if err != nil {
		err = fmt.Errorf("characterHandler.GetAll()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ObjectsDTO := make([]dto.CharacterDTO, len(Objects))
	for i := range Objects {
		ObjectsDTO[i] = mapper.ToCharacterDTO(*Objects[i])
	}
	c.JSON(http.StatusOK, ObjectsDTO)
}

// DELETE object
func (h *characterHandler) Delete(c *gin.Context) {
	id := c.GetHeader("id")
	if id == "" {
		err := errors.New("characterHandler.Delete(): No ID header in request")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := h.Usecase.Delete(c.Request.Context(), entity.ID(id))
	if err != nil {
		err = fmt.Errorf("characterHandler.Delete()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "DELETE: SUCCESS"})
}

// PUT object
func (h *characterHandler) Set(c *gin.Context) {
	var DTO dto.CharacterDTO
	if err := c.ShouldBindJSON(&DTO); err != nil {
		err = fmt.Errorf("characterHandler.Set()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный JSON"})
		return
	}
	object := mapper.ToCharacterEntity(DTO)
	clientID, _ := c.Get("client_id")
	object.CreatorID = clientID.(entity.ID).String()
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
