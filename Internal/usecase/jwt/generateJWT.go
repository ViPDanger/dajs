package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var AccessTokenTime = time.Minute
var RefreshTokenTime = time.Hour * 24 * 7

var AccessKey = []byte("ACCESSKEY")
var RefreshKey = []byte("REFRESHKEY")

func NewAccessToken(username string) (accessToken string, exp time.Time, err error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(AccessTokenTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString(AccessKey)
	if err != nil {
		exp = time.Now().Add(AccessTokenTime)
	}
	return
}

func NewRefreshToken(username string) (refreshToken string, exp time.Time, err error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(RefreshTokenTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err = token.SignedString(RefreshKey)
	if err != nil {
		exp = time.Now().Add(RefreshTokenTime)
	}
	return
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

func RefreshAccessToken(refreshTokenString string) (accessToken string, exp time.Time, err error) {
	c := make(jwt.MapClaims)
	token, err := jwt.ParseWithClaims(refreshTokenString, c, func(token *jwt.Token) (interface{}, error) {
		return RefreshKey, nil
	})
	if err != nil {
		return
	}
	if !token.Valid {
		err = errors.New("Invalid Token")
		return
	}

	claims := jwt.MapClaims{
		"username": c["username"],
		"exp":      time.Now().Add(AccessTokenTime).Unix(),
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString(AccessKey)
	if err != nil {
		exp = time.Now().Add(AccessTokenTime)
	}
	return
}
