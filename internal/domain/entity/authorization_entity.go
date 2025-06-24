package entity

import "time"

type AccessToken struct {
	Str        string
	ExpireTime time.Time
}
type RefreshToken struct {
	Str        string
	ExpireTime time.Time
}
type User struct {
	Uid      string
	Name     string
	Password string
}

func (c User) GetID() string {
	return c.Name
}
