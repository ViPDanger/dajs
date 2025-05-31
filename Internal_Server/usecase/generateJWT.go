package usecase

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var AccessTokenTime = time.Second
var RefreshTokenTime = time.Hour * 24 * 7 //100 * time.Millisecond //

var AccessKey = []byte("ACCESSKEY")
var RefreshKey = []byte("REFRESHKEY")

func NewAccessToken(username string) (string, error) {
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

func ParseAccessToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return AccessKey, nil
	})
}
func ParseRefreshToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return RefreshKey, nil
	})
}

func RefreshAccessToken(refreshTokenString string) (string, error) {
	c := make(jwt.MapClaims)
	token, err := jwt.ParseWithClaims(refreshTokenString, c, func(token *jwt.Token) (interface{}, error) {
		return RefreshKey, nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("Invalid Token")
	}

	claims := jwt.MapClaims{
		"username": c["username"],
		"exp":      time.Now().Add(AccessTokenTime).Unix(),
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(AccessKey)
}
