package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	dto "github.com/ViPDanger/dajs/go-api/internal/interfaces/dto"
	"github.com/ViPDanger/dajs/go-api/internal/interfaces/jwt"
	"github.com/ViPDanger/dajs/go-api/internal/interfaces/mapper"
	"github.com/ViPDanger/dajs/go-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Registration(c *gin.Context)
	Login(c *gin.Context)
	Refresh(c *gin.Context)
}

func NewUserHandler(uc usecase.UserUseCase) UserHandler {
	return &userHandler{UserUseCase: uc}
}

// ИСПРАВИТЬ!!!
type userHandler struct {
	usecase.UserUseCase
}

// REGISTER USER -----------
func (userHandler *userHandler) Registration(c *gin.Context) {
	var userDTO dto.UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		err = fmt.Errorf("userHandler.Registration()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "UnmarshalJSON error"})
		return
	}
	var err error
	var id *entity.ID
	if id, err = userHandler.UserUseCase.Register(c.Request.Context(), mapper.ToUserEntity(userDTO)); err != nil && id != nil {
		err = fmt.Errorf("userHandler.Registration()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusCreated, gin.H{"id": *id})
}

// LOGIN USER	--------
func (userHandler *userHandler) Login(c *gin.Context) {
	var userDTO dto.UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "UnmarshalJSON error"})
		return
	}
	user := mapper.ToUserEntity(userDTO)
	err := userHandler.UserUseCase.Login(c.Request.Context(), user)
	if err != nil {
		err = fmt.Errorf("userHandler.Login()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	accessToken, err := jwt.NewAccessToken(user.Name)
	if err != nil {
		err = fmt.Errorf("userHandler.Login()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	refreshToken, err := jwt.NewRefreshToken(user.Name)
	if err != nil {
		err = fmt.Errorf("userHandler.Login()/%w", err)
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
func (userHandler *userHandler) Refresh(c *gin.Context) {
	var refreshDTO dto.RefreshTokenDTO
	if err := c.ShouldBindJSON(&refreshDTO); err != nil {
		err = fmt.Errorf("userHandler.Refresh()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	refreshToken := mapper.ToRefreshTokenEntity(refreshDTO)
	token, err := jwt.ParseRefreshToken(refreshToken.Str)
	if err != nil || !token.Valid {
		err = errors.New("Invalid token")
		err = fmt.Errorf("userHandler.Refresh()/%w", err)
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	accessToken, err := jwt.RefreshAccessToken(refreshToken.Str)
	if err != nil {
		_ = c.Error(err)
		err = fmt.Errorf("userHandler.Refresh()/%w", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, mapper.ToAccessTokenDTO(accessToken))
}

// MiddleWare access token check
func Protected(c *gin.Context) {
	accessHeader := c.GetHeader("Authorization")
	if accessHeader == "" {
		err := errors.New("userHandler.Protected():No token in protected request")
		_ = c.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен отсутствует"})
		return
	}
	tokenString := strings.TrimPrefix(accessHeader, "Bearer ")
	token, err := jwt.ParseAccessToken(tokenString)
	if err != nil || !token.Valid {
		err = errors.New("userHandler.Protected(): Invalid token")
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return

	}
	c.Next()
}
