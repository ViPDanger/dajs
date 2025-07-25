package jwt

import (
	"errors"
	"time"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"

	"github.com/golang-jwt/jwt/v5"
)

var AccessTokenTime = time.Minute
var RefreshTokenTime = time.Hour * 24 * 7

var AccessKey = []byte("ACCESSKEY")
var RefreshKey = []byte("REFRESHKEY")

func NewAccessToken(id string) (accessToken entity.AccessToken, err error) {

	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(AccessTokenTime).Unix(),
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

func NewRefreshToken(id string) (refreshToken entity.RefreshToken, err error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(RefreshTokenTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshString, err := token.SignedString(RefreshKey)
	var expTime time.Time
	if err != nil {
		expTime = time.Now().Add(RefreshTokenTime)
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
	return NewAccessToken(c["id"].(string))
}

func GetIDFromToken(token *jwt.Token) (id *entity.ID, err error) {
	if mapClaims, ok := token.Claims.(jwt.MapClaims); ok {
		if s, ok := mapClaims["id"].(string); ok {
			id := entity.ID(s)
			return &id, nil
		}
		return nil, errors.New("GetIdFromToken(): no id in mapClaims")
	}
	return nil, errors.New("GetIdFromToken(): can't parse mapClaims from token")
}
