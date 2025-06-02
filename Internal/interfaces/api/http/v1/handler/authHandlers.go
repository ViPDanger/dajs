package handler

import (
	dto "DAJ/Internal/interfaces/api/dto"
	"DAJ/Internal/usecase/jwt"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ИСПРАВИТЬ!!!
var users = map[string]string{} // Имитация БД: username -> hashedPassword

// REGISTER USER -----------
func Register(c *gin.Context) {
	var request dto.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, dto.Error{Error: "Некорректный JSON"})
		return
	}
	//	ИСПРАВИИИИИИИИИИТЬ
	if _, exists := users[request.Username]; exists {
		_ = c.Error(errors.New("user currently existing"))
		c.JSON(http.StatusConflict, dto.Error{Error: "Пользователь уже существует"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.Error{Error: "Ошибка шифрования"})
		return
	}

	users[request.Username] = string(hash)
	c.JSON(http.StatusCreated, dto.Message{
		Message: "Регистрация прошла успешно",
	})
}

// LOGIN USER	--------
func Login(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, dto.Error{Error: "Некорректный JSON"})
		return
	}
	// АААААААААААААААААААААААААААААААА МЕНЯЙ ТВАРЬ
	storedPassword, exists := users[request.Username]
	if !exists || bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(request.Password)) != nil {
		_ = c.Error(errors.New("Username or password is incorrect"))
		c.JSON(http.StatusUnauthorized, dto.Error{Error: "Неверные имя пользователя или пароль"})
		return
	}

	accessToken, accessTokenExpireTime, err := jwt.NewAccessToken(request.Username)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.Error{Error: "Не удалось создать токен"})
		return
	}
	refreshToken, refreshTokenExpireTime, err := jwt.NewRefreshToken(request.Username)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.Error{Error: "Не удалось создать токен"})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		AccessToken:    accessToken,
		AccessExpTime:  accessTokenExpireTime,
		RefreshToken:   refreshToken,
		RefreshExpTime: refreshTokenExpireTime,
	})
}

// REFRESH ACCESS TOKEN by REFRESH TOKEN
func Refresh(c *gin.Context) {
	var request dto.RefreshRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, dto.Error{Error: "Некорректный JSON"})
		return
	}
	token, err := jwt.ParseRefreshToken(request.RefreshToken)
	if err != nil || !token.Valid {
		_ = c.Error(errors.New("Invalid token"))
		c.JSON(http.StatusUnauthorized, dto.Error{Error: "Невалидный токен"})
		return
	}
	accessToken, accessTokenExpireTime, err := jwt.RefreshAccessToken(request.RefreshToken)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, dto.Error{Error: "Не удалось создать токен"})
		return
	}
	c.JSON(http.StatusOK, dto.RefreshResponse{
		AccessToken:   accessToken,
		AccessExpTime: accessTokenExpireTime,
	})
}
