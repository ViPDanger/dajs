package models

import "time"

type AccessToken struct {
	AccessString string    `json:"access_token"`
	AccessExp    time.Time `json:"access_exp"`
}
type RefreshToken struct {
	RefreshString string    `json:"refresh_token"`
	RefreshExp    time.Time `json:"refresh_exp"`
}

type User struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Message
}
type RefreshResponse struct {
	AccessToken
}

type AuthResponse struct {
	AccessToken
	RefreshToken
}

type ProtectedResponse struct {
	Message
}
