package v1_01

import (
	"DAJ/Internal_Server/models"
	"DAJ/Internal_Server/usecase"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

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

	token, err := usecase.NewAccesToken(user.Username)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать токен"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Protected(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		_ = c.Error(errors.New("No token in protected request"))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен отсутствует"})
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := usecase.ParseToken(tokenString)
	if err != nil || !token.Valid {
		_ = c.Error(errors.New("Invalid token"))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Невалидный токен"})
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	c.JSON(http.StatusOK, gin.H{"message": "Добро пожаловать, " + username})
}
