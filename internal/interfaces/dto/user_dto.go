package dto

import "time"

type UserDTO struct {
	Username string `json:"user"`
	Password string `json:"password"`
}
type AccessTokenDTO struct {
	Str        string    `json:"access_token"`
	ExpireTime time.Time `json:"access_exp"`
}
type RefreshTokenDTO struct {
	Str        string    `json:"refresh_token"`
	ExpireTime time.Time `json:"refresh_exp"`
}
type TokensDTO struct {
	Access  AccessTokenDTO
	Refresh RefreshTokenDTO
}
