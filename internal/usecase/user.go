package usecase

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	Repo repository.Repository[entity.User]
}

func (UC *UserUseCase) Register(user entity.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	if err := UC.Repo.New(user.Name, &user); err != nil {
		return err
	}

	return nil
}
func (UC *UserUseCase) Login(user entity.User) error {
	if repoUser, err := UC.Repo.GetByID(user.Name); err != nil {
		return err
	} else if bcrypt.CompareHashAndPassword([]byte(repoUser.Password), []byte(user.Password)) != nil {
		return errors.New("Неверные имя пользователя или пароль")
	}

	return nil
}

func (UC *UserUseCase) Delete(character entity.User) error {
	return UC.Repo.Delete(character.Name)
}
