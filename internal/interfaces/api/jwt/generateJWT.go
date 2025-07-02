package jwt

import (
	"errors"
	"time"

	"github.com/ViPDanger/dajs/internal/domain/entity"

	"github.com/golang-jwt/jwt/v5"
)

var AccessTokenTime = time.Minute
var RefreshTokenTime = time.Hour * 24 * 7

var AccessKey = []byte("ACCESSKEY")
var RefreshKey = []byte("REFRESHKEY")

func NewAccessToken(username string) (accessToken entity.AccessToken, err error) {

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(AccessTokenTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessString, err := token.SignedString(AccessKey)
	var expTime time.Time
	if err != nil {
		expTime = time.Now().Add(AccessTokenTime)
	}

	return entity.AccessToken{
		Str:        accessString,
		ExpireTime: expTime,
	}, err
}

func NewRefreshToken(username string) (refreshToken entity.RefreshToken, err error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(RefreshTokenTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshString, err := token.SignedString(AccessKey)
	var expTime time.Time
	if err != nil {
		expTime = time.Now().Add(AccessTokenTime)
	}

	return entity.RefreshToken{
		Str:        refreshString,
		ExpireTime: expTime,
	}, err
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

func RefreshAccessToken(refreshTokenString string) (accessToken entity.AccessToken, err error) {
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
	accessString, err := token.SignedString(AccessKey)
	var expTime time.Time
	if err != nil {
		expTime = time.Now().Add(AccessTokenTime)
	}

	return entity.AccessToken{
		Str:        accessString,
		ExpireTime: expTime,
	}, err
}
