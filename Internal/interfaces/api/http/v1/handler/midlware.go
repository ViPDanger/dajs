package handler

import (
	usecase "DAJ/Internal/usecase/jwt"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// MiddleWare access token check
func Protected(c *gin.Context) {
	accessHeader := c.GetHeader("Authorization")
	if accessHeader == "" {
		_ = c.Error(errors.New("No token in protected request"))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен отсутствует"})
		return
	}
	tokenString := strings.TrimPrefix(accessHeader, "Bearer ")
	token, err := usecase.ParseAccessToken(tokenString)
	if err != nil || !token.Valid {
		_ = c.Error(errors.New("Invalid token"))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Невалидный токен"})
		return
	}
	c.Next()
}
