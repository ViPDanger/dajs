package entity

import "errors"

type User struct {
	Uid      string
	Name     string
	password string
}

func (u *User) IsPassword(password string) error {
	if u.password != password {
		return errors.New("password is incorrect")
	}
	return nil
}
func (u *User) SetPassword(oldPassword, newPassword string) (err error) {
	err = u.IsPassword(oldPassword)
	if err==nil{
		u.password = newPassword
	}
	return
}