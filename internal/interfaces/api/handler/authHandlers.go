package handler

import (
	dto "DAJ/internal/interfaces/api/dto"
	"DAJ/internal/interfaces/api/jwt"
	"DAJ/internal/interfaces/api/mapper"
	"DAJ/internal/usecase"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ИСПРАВИТЬ!!!
type UserHandler struct {
	UserUC usecase.UserUseCase
}

// REGISTER USER -----------
func (userHandler *UserHandler) Registration(c *gin.Context) {
	var userDTO dto.UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, dto.Error{Error: "UnmarshalJSON error"})
		return
	}
	if err := userHandler.UserUC.Register(mapper.ToUserEntity(userDTO)); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusCreated, dto.Message{
		Message: "Registration: Success",
	})
}

// LOGIN USER	--------
func (userHandler *UserHandler) Login(c *gin.Context) {
	var userDTO dto.UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, dto.Error{Error: "UnmarshalJSON error"})
		return
	}
	user := mapper.ToUserEntity(userDTO)
	err := userHandler.UserUC.Login(user)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	accessToken, err := jwt.NewAccessToken(user.Name)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	refreshToken, err := jwt.NewRefreshToken(user.Name)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, dto.TokensDTO{
		Access:  mapper.ToAccessTokenDTO(accessToken),
		Refresh: mapper.ToRefreshTokenDTO(refreshToken),
	})
}

// REFRESH ACCESS TOKEN by REFRESH TOKEN
func (userHandler *UserHandler) Refresh(c *gin.Context) {
	var refreshDTO dto.RefreshTokenDTO
	if err := c.ShouldBindJSON(&refreshDTO); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	refreshToken := mapper.ToRefreshTokenEntity(refreshDTO)
	token, err := jwt.ParseRefreshToken(refreshToken.Str)
	if err != nil || !token.Valid {
		err = errors.New("Invalid token")
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	accessToken, err := jwt.RefreshAccessToken(refreshToken.Str)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, mapper.ToAccessTokenDTO(accessToken))
}

// MiddleWare access token check
func Protected(c *gin.Context) {
	accessHeader := c.GetHeader("Authorization")
	if accessHeader == "" {
		_ = c.Error(errors.New("No token in protected request"))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен отсутствует"})
		return
	}
	tokenString := strings.TrimPrefix(accessHeader, "Bearer ")
	token, err := jwt.ParseAccessToken(tokenString)
	if err != nil || !token.Valid {
		_ = c.Error(errors.New("Invalid token"))
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Невалидный токен"})
		return

	}
	c.Next()
}
