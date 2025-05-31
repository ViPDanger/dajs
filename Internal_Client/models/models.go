package models

import (
	"time"
)

type accessToken struct {
	AccessToken string        `json:"access_token"`
	AccessExp   time.Duration `json:"access_exp"`
}
type refreshToken struct {
	RefreshToken string        `json:"refresh_token"`
	RefreshExp   time.Duration `json:"refresh_exp"`
}

type User struct {
	Username string `json:"user"`
	Password string `json:"password"`
}
type Message struct {
	Message string `json:"message"`
}

type RegisterResponse struct {
	Message
}
type RefreshResponse struct {
	accessToken
}

type AuthResponse struct {
	accessToken
	refreshToken
}
