package usecase

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var AccessTokenTime = time.Hour
var RefreshTokenTime = time.Hour * 24 * 7

var AccessKey = []byte("ACCESSKEY")
var RefreshKey = []byte("ACCESSKEY")

func NewAccesToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(AccessTokenTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(AccessKey)
}

func NewRefreshToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(RefreshTokenTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(RefreshKey)
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return AccessKey, nil
	})
}
