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
	Usecase usecase.MonsterUsecase
}

func NewMonsterHandler(uc usecase.MonsterUsecase) *monsterHandler {
	return &monsterHandler{Usecase: uc}
}

// GET object
func (h *monsterHandler) Get(c *gin.Context) {
	// проверка id header
	id := c.GetHeader("id")
	if id == "" {
		err := errors.New("monsterHandler.Get():No id header in request")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// обращение к Usecase
	object, err := h.Usecase.GetByID(c.Request.Context(), entity.ID(id))
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
	// проверка clientId header
	clientId, ok := c.Get("client_id")
	if !ok {
		err := errors.New("monsterHandler.Get(): client_id не найден")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// обращение к Usecase
	objects, err := h.Usecase.GetByCreatorID(c.Request.Context(), clientId.(entity.ID))
	if err != nil {
		err = fmt.Errorf("monsterHandler.Get()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dtos := make([]dto.MonsterDTO, len(objects))

	for i := range objects {
		dtos[i] = mapper.ToMonsterDTO((objects[i]))
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
	object.CreatorID = clientID.(entity.ID)
	// Обращение к Usecase
	id, err := h.Usecase.New(c.Request.Context(), &object)
	if err != nil || id == nil {
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
	Objects, err := h.Usecase.GetAll(c.Request.Context())
	if err != nil {
		err = fmt.Errorf("monsterHandler.GetAll()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ObjectsDTO := make([]dto.MonsterDTO, len(Objects))
	for i := range Objects {
		ObjectsDTO[i] = mapper.ToMonsterDTO(Objects[i])
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
	err := h.Usecase.Delete(c.Request.Context(), entity.ID(id))
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
	clientID, _ := c.Get("client_id")
	object.CreatorID = clientID.(entity.ID)
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
