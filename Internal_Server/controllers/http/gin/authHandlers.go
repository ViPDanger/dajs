package gin

import (
	"DAJ/Internal_Server/models"
	"DAJ/Internal_Server/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var accessTokenExpireTime = usecase.AccessTokenTime
var refreshTokenExpireTime = usecase.RefreshTokenTime

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный JSON"})
		return
	}

	if _, exists := users[user.Username]; exists {
		_ = c.Error(errors.New("user currently existing"))
		c.JSON(http.StatusConflict, gin.H{"error": "Пользователь уже существует"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка шифрования"})
		return
	}

	users[user.Username] = string(hash)
	c.JSON(http.StatusCreated, gin.H{"message": "Регистрация прошла успешно"})
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный JSON"})
		return
	}

	storedPassword, exists := users[user.Username]
	if !exists || bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password)) != nil {
		_ = c.Error(errors.New("Username or password is incorrect"))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные имя пользователя или пароль"})
		return
	}

	access_token, err := usecase.NewAccessToken(user.Username)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать токен"})
		return
	}
	refresh_token, err := usecase.NewRefreshToken(user.Username)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать токен"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": access_token, "access_exp": accessTokenExpireTime, "refresh_token": refresh_token, "refresh_exp": refreshTokenExpireTime})
}

func Refresh(c *gin.Context) {
	var refreshToken models.RefreshToken
	if err := c.ShouldBindJSON(&refreshToken); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный JSON"})
		return
	}
	token, err := usecase.ParseRefreshToken(refreshToken.RefreshToken)
	if err != nil || !token.Valid {
		_ = c.Error(errors.New("Invalid token"))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Невалидный токен"})
		return
	}
	access_token, err := usecase.RefreshAccessToken(refreshToken.RefreshToken)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать токен"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": access_token, "expires_in": accessTokenExpireTime})
}
