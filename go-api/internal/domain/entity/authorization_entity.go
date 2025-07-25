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
	ID       `bson:"_id"`
	Username string
	Password string
}
