package entity

import "time"

type AccessToken struct {
	Str        string    `json:"access_token"`
	ExpireTime time.Time `json:"access_exp"`
}
type RefreshToken struct {
	Str        string    `json:"refresh_token"`
	ExpireTime time.Time `json:"refresh_exp"`
}
type User struct {
	Uid      string
	Name     string
	Password string
}

func (c User) GetID() string {
	return c.Name
}
