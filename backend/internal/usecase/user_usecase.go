package usecase

import (
	"backend/internal/repository"
	"backend/internal/utils"
	"context"
	"errors"
	"fmt"
	"os"
)

var secretKey = os.Getenv("JWT_SECRET")

type UserUsecase struct {
	repo *repository.UserRepository
}

func NewUserUsecase(repo *repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (u *UserUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.repo.GetUserByEmail(context.Background(), email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !utils.ComparePassword(user.Password, password) {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(int(user.ID), secretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (u *UserUsecase) Register(ctx context.Context, username, email, password string) error {
	emailExists, err := u.repo.EmailExists(ctx, email)
	if err != nil {
		return err
	}
	if emailExists {
		return errors.New("email already exists")
	}

	usernameExists, err := u.repo.UsernameExists(ctx, username)
	if err != nil {
		return err
	}
	if usernameExists {
		return errors.New("username already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Printf("error hashing password:%v\n", err)
		return err
	}

	err = u.repo.CreateUser(ctx, username, email, hashedPassword)
	if err != nil {
		fmt.Printf("error creating user:%v\n", err)
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}
