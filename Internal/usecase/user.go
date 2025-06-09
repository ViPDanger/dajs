package usecase

import (
	"DAJ/Internal/domain/entity"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var users = map[string]string{}

type UserUseCase struct {
}

func (userUC *UserUseCase) Register(character entity.User) error {
	if _, exists := users[character.Name]; exists {
		return errors.New("Пользователь уже существует")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(character.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	users[character.Name] = string(hash)
	return nil
}
func (userUC *UserUseCase) Login(character entity.User) (err error) {
	storedPassword, exists := users[character.Name]
	if !exists || bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(character.Password)) != nil {
		err = errors.New("Неверные имя пользователя или пароль")
		return
	}

	return nil
}

func (userUC *UserUseCase) Delete(character entity.User) error {
	return nil
}
