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

type monsterHandler struct {
	UC usecase.MonsterUsecase
}

func NewMonsterHandler(UC usecase.MonsterUsecase) *monsterHandler {
	return &monsterHandler{}
}

// GET object
func (h *monsterHandler) Get(c *gin.Context) {
	// проверка id header
	id := c.GetHeader("id")
	if id == "" {
		err := errors.New("monsterHandler.Get(): Нет ID в запросе")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// обращение к Usecase
	object, err := h.UC.GetByID(c.Request.Context(), entity.ID(id))
	if err != nil {
		err = fmt.Errorf("monsterHandler.Get()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// ВЫВОД
	c.JSON(http.StatusOK, mapper.ToMonsterDTO(*object))
}

func (h *monsterHandler) GetByCreatorID(c *gin.Context) {
	// проверка id header
	clientId, ok := c.Get("client_id")
	if !ok {
		err := errors.New("monsterHandler.Get(): Нет ID в запросе")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// обращение к Usecase

	objects, err := h.UC.GetByCreatorID(c.Request.Context(), entity.ID(clientId.(string)))
	if err != nil {
		err = fmt.Errorf("monsterHandler.Get()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dtos := make([]dto.MonsterDTO, len(objects))
	for i := range objects {
		dtos[i] = mapper.ToMonsterDTO(*(objects[i]))
	}
	// ВЫВОД
	c.JSON(http.StatusOK, dtos)
}

// POST object
func (h *monsterHandler) New(c *gin.Context) {
	var DTO dto.MonsterDTO
	// проверка body запроса
	if err := c.ShouldBindJSON(&DTO); err != nil {
		err = fmt.Errorf("monsterHandler.New()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "UnmarshalJSON error"})
		return
	}
	object := mapper.ToMonsterEntity(DTO)
	clientID, _ := c.Get("client_id")
	object.CreatorID = clientID.(entity.ID).String()
	// Обращение к Usecase
	id, err := h.UC.New(c.Request.Context(), &object)
	if err != nil || id != nil {
		err = fmt.Errorf("monsterHandler.New()/%w", err)
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
func (h *monsterHandler) GetAll(c *gin.Context) {
	Objects, err := h.UC.GetAll(c.Request.Context())
	if err != nil {
		err = fmt.Errorf("monsterHandler.GetAll()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ObjectsDTO := make([]dto.MonsterDTO, len(Objects))
	for i := range Objects {
		ObjectsDTO[i] = mapper.ToMonsterDTO(*Objects[i])
	}
	c.JSON(http.StatusOK, ObjectsDTO)
}

// DELETE object
func (h *monsterHandler) Delete(c *gin.Context) {
	id := c.GetHeader("id")
	if id == "" {
		err := errors.New("monsterHandler.Delete(): No ID header in request")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := h.UC.Delete(c.Request.Context(), entity.ID(id))
	if err != nil {
		err = fmt.Errorf("monsterHandler.Delete()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "DELETE: SUCCESS"})
}

// PUT object
func (h *monsterHandler) Set(c *gin.Context) {
	var DTO dto.MonsterDTO
	if err := c.ShouldBindJSON(&DTO); err != nil {
		err = fmt.Errorf("monsterHandler.Set()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный JSON"})
		return
	}
	object := mapper.ToMonsterEntity(DTO)
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
