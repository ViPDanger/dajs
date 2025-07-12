package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/domain/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Register(ctx context.Context, user entity.User) (*entity.ID, error)
	Login(ctx context.Context, user entity.User) error
	Delete(ctx context.Context, character entity.ID) error
}

type userUseCase struct {
	repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{UserRepository: repo}
}

func (uc *userUseCase) Register(ctx context.Context, user entity.User) (id *entity.ID, err error) {
	if uc.UserRepository == nil {
		return nil, errors.New("userUseCase.Register(): Nill pointer repository")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("userUseCase.Register()/%w", err)
	}
	user.Password = string(hash)
	if id, err = uc.UserRepository.Insert(ctx, &user); err != nil {
		return nil, fmt.Errorf("userUseCase.Register()/%w", err)
	}
	return
}
func (uc *userUseCase) Login(ctx context.Context, user entity.User) (err error) {
	if uc.UserRepository == nil {
		return errors.New("userUseCase.Login(): Nill pointer repository")
	}
	if repoUser, err := uc.UserRepository.GetByID(ctx, user.GetID()); err != nil {
		return fmt.Errorf("userUseCase.Login()/%w", err)
	} else if bcrypt.CompareHashAndPassword([]byte(repoUser.Password), []byte(user.Password)) != nil {
		return fmt.Errorf("userUseCase.Login(): Wrong login or password")
	}
	return nil
}
func (uc *userUseCase) Delete(ctx context.Context, id entity.ID) error {
	if uc.UserRepository == nil {
		return errors.New("userUseCase.Delete(): Nill pointer repository")
	}
	if err := uc.UserRepository.Delete(ctx, id); err != nil {
		return fmt.Errorf("userUseCase.Delete()/%w", err)
	}
	return nil
}
