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

type characterHandler struct {
	Usecase usecase.CharacterUsecase
}

func NewCharacterHandler(uc usecase.CharacterUsecase) *characterHandler {
	return &characterHandler{Usecase: uc}
}

// GET object

func (h *characterHandler) Get(c *gin.Context) {
	// проверка clientId header
	var creator_id string
	if clientId, ok := c.Get("client_id"); !ok {
		err := errors.New("characterHandler.Get(): client_id не найден")
		_ = c.Error(err)
		//c.JSON(http.StatusBadRequest, err)
		//return
	} else {
		creator_id = clientId.(string)
	}
	var ids []string
	for _, s := range c.QueryArray("id") {
		ids = append(ids, string(s))
	}

	// обращение к Usecase
	objects, err := h.Usecase.Get(c.Request.Context(), creator_id, ids...)
	if err != nil {
		err = fmt.Errorf("characterHandler.Get()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(objects) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"error": "No Content"})
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
	object.CreatorID, _ = clientID.(string)
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
		"id": id,
	})
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
	err := h.Usecase.Delete(c.Request.Context(), string(id))
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
