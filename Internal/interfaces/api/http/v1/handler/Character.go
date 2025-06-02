package handler

import (
	"DAJ/Internal/interfaces/api/dto"
	"DAJ/Internal/interfaces/mapper"
	"DAJ/Internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CharacterHandler struct {
	characterUC usecase.CharacterUseCase
}

func New(characterUC usecase.CharacterUseCase) CharacterHandler {
	return CharacterHandler{characterUC: characterUC}
}

func (ch *CharacterHandler) GetCharacter(c *gin.Context) {
	id := c.GetHeader("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, errors.New("Нет ID в запросе"))
	}
	char, err := ch.characterUC.Get(id)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, mapper.CharacterToCharacterDTO(char))
}

func (ch *CharacterHandler) NewCharacter(c *gin.Context) {
	var DTO dto.CharacterDTO
	if err := c.ShouldBindJSON(&DTO); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, dto.Error{Error: "Некорректный JSON"})
		return
	}
	err := ch.characterUC.New(mapper.CharacterDTOtoCharacter(DTO))
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.Message{
		Message: "Character with id" + DTO.ID + " created",
	})
}
