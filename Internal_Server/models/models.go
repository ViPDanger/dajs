package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"` // Хранится в зашифрованном виде
}
type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}
