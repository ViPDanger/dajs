package dto

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type AccessTokenDTO struct {
	AccessToken string `json:"access_token"`
}
type RefreshTokenDTO struct {
	RefreshToken string `json:"refresh_token"`
}

type TokensDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
