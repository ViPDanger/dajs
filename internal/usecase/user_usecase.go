package usecase

import (
	"fmt"

	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/domain/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	Repo repository.Repository[entity.User]
}

func NewUserUsecase(repo repository.Repository[entity.User]) *UserUseCase {
	return &UserUseCase{Repo: repo}
}

func (UC *UserUseCase) Register(user entity.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("UserUseCase.Register()/%w", err)
	}
	user.Password = string(hash)
	if err := UC.Repo.Insert(&user); err != nil {
		return fmt.Errorf("UserUseCase.Register()/%w", err)
	}

	return nil
}
func (UC *UserUseCase) Login(user entity.User) error {
	if repoUser, err := UC.Repo.GetByID(user.Name); err != nil {
		return fmt.Errorf("UserUseCase.Login()/%w", err)
	} else if bcrypt.CompareHashAndPassword([]byte(repoUser.Password), []byte(user.Password)) != nil {
		return fmt.Errorf("UserUseCase.Login(): Wrong login or password")
	}

	return nil
}

func (UC *UserUseCase) Delete(character entity.User) error {
	if err := UC.Repo.Delete(character.Name); err != nil {
		return fmt.Errorf("UserUseCase.Delete()/%w", err)
	}
	return nil
}
