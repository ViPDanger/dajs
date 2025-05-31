package gin

import (
	"DAJ/Internal_Server/usecase"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ProtectedHandleFunc(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		_ = c.Error(errors.New("No token in protected request"))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен отсутствует"})
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := usecase.ParseAccessToken(tokenString)
	if err != nil || !token.Valid {
		_ = c.Error(errors.New("Invalid token"))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Невалидный токен"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	c.JSON(http.StatusOK, gin.H{"message": "Добро пожаловать, " + username})
	c.Next()
}
