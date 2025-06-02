package dto

import "time"

// Registration
type RegisterRequest struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

// Login
type LoginResponse struct {
	AccessToken    string    `json:"access_token"`
	AccessExpTime  time.Time `json:"access_exp"`
	RefreshToken   string    `json:"refresh_token"`
	RefreshExpTime time.Time `json:"refresh_exp"`
}

// Refresh
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken   string    `json:"access_token"`
	AccessExpTime time.Time `json:"access_exp"`
}
