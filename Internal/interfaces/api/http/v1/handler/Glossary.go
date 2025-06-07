package handler

import (
	"DAJ/Internal/interfaces/api/dto"
	"DAJ/Internal/interfaces/api/mapper"
	"DAJ/Internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GlossaryHandler struct {
	GlossaryUC usecase.GlossaryUseCase
}

func NewGlossaryHandler(GlossaryUC usecase.GlossaryUseCase) GlossaryHandler {
	return GlossaryHandler{GlossaryUC: GlossaryUC}
}

// GET Glossary
func (ch *GlossaryHandler) GetGlossary(c *gin.Context) {
	id := c.GetHeader("id")
	if id == "" {
		err := errors.New("Нет ID в запросе")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	char, err := ch.GlossaryUC.Get(id)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, mapper.ToGlossaryDTO(char))
}

// POST Glossary
func (ch *GlossaryHandler) NewGlossary(c *gin.Context) {
	var DTO dto.GlossaryDTO
	if err := c.ShouldBindJSON(&DTO); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, dto.Error{Error: "Некорректный JSON"})
		return
	}
	err := ch.GlossaryUC.New(mapper.ToGlossaryEntity(DTO))
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.Message{
		Message: "Glossary with id " + DTO.ID + " created",
	})
}

// GET all chatacter
func (ch *GlossaryHandler) GetAllGlossarys(c *gin.Context) {
	Glossarys, err := ch.GlossaryUC.GetAll()
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
		return
	}
	GlossarysDTO := make([]dto.GlossaryDTO, len(Glossarys))
	for i := range Glossarys {
		GlossarysDTO = append(GlossarysDTO, mapper.ToGlossaryDTO(Glossarys[i]))
	}
	c.JSON(http.StatusOK, GlossarysDTO)
}

// DELETE Glossary
func (ch *GlossaryHandler) DeleteGlossary(c *gin.Context) {
	id := c.GetHeader("id")
	if id == "" {
		err := errors.New("Нет ID в запросе")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := ch.GlossaryUC.Delete(id)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, dto.Message{Message: "Glossary with id " + id + "deleted"})
}

// PUT Glossary
func (ch *GlossaryHandler) SetGlossary(c *gin.Context) {
	var DTO dto.GlossaryDTO
	if err := c.ShouldBindJSON(&DTO); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, dto.Error{Error: "Некорректный JSON"})
		return
	}
	err := ch.GlossaryUC.Set(mapper.ToGlossaryEntity(DTO))
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.Message{
		Message: "Успешно",
	})
}
