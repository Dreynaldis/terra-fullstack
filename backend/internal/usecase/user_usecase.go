package usecase

import (
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/utils"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
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

func (u *UserUsecase) Login(ctx context.Context, loginInput, password string) (string, *model.User, error) {
	var user *model.User
	var err error

	if strings.Contains(loginInput, "@") {
		user, err = u.repo.GetUserByEmail(ctx, loginInput)
	} else {
		user, err = u.repo.GetUserByUsername(ctx, loginInput)
	}

	if err != nil {
		return "", nil, errors.New("user not found")
	}
	if user.Provider == "local" && !utils.ComparePassword(user.Password, password) {
		return "", nil, errors.New("invalid password")
	}
	if user.Provider == "google" {
		return "", nil, errors.New("this account is registered with google. please login with google")
	}

	token, err := utils.GenerateJWT(int(user.ID), secretKey)
	if err != nil {
		return "", nil, err
	}
	fmt.Printf("token: %v\n", token)

	return token, user, nil
}
func (u *UserUsecase) Register(ctx context.Context, username, email, password, provider string) error {

	if provider == "google" {
		existingUser, err := u.repo.GetUserByEmail(ctx, email)
		if err == nil && existingUser.Provider != "google" {
			return u.repo.UpdateUserProvider(ctx, email, "google", password)
		} else if err != nil && err.Error() == "usernot found" {
			return u.repo.CreateUser(ctx, username, email, password, "google")
		}
	}
	if strings.Contains(email, "@") || strings.Contains(username, ".") {
		return errors.New("username can't have invalid characters")
	}
	if !utils.IsValidEmail(email) {
		return errors.New("invalid email")
	}
	emailExists, err := u.repo.EmailExists(ctx, email)
	if err != nil {
		return fmt.Errorf("error checking email existence: %w", err)
	}
	if emailExists {
		return errors.New("email already exists")
	}
	usernameExists, err := u.repo.UsernameExists(ctx, username)
	if err != nil {
		return fmt.Errorf("error checking username existence: %w", err)
	}
	if usernameExists {
		return errors.New("username already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	err = u.repo.CreateUser(ctx, username, email, hashedPassword, "local")
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func (u *UserUsecase) HandleGoogleCallback(ctx context.Context, googleToken string) (string, string, error) {

	profile, err := utils.FetchGoogleProfile(googleToken)
	if err != nil {
		return "", "", fmt.Errorf("failed to fetch google profile: %w", err)
	}

	email := profile.Email
	username := profile.Name
	if username == "" {
		username = profile.GivenName
	}

	emailExists, err := u.repo.EmailExists(ctx, email)
	if err != nil {
		return "", "", fmt.Errorf("failed to check email existence: %w", err)
	}
	if !emailExists {
		hashedPassword, err := utils.HashPassword(profile.Sub)
		if err != nil {
			return "", "", fmt.Errorf("failed to hash password: %w", err)
		}
		err = u.repo.CreateUser(ctx, username, email, hashedPassword, "google")
		if err != nil {
			return "", "", fmt.Errorf("failed to create user: %w", err)
		}
	} else {
		hashedPassword, err := utils.HashPassword(profile.Sub)
		if err != nil {
			return "", "", fmt.Errorf("failed to hash password: %w", err)
		}
		//change user provider to google if existing user has local provider
		err = u.repo.UpdateUserProvider(ctx, email, "google", hashedPassword)
		if err != nil {
			return "", "", fmt.Errorf("failed to update user provider: %w", err)
		}
	}

	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", "", fmt.Errorf("failed to get user: %w", err)
	}

	token, err := utils.GenerateJWT(int(user.ID), secretKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate JWT: %w", err)
	}
	return token, user.Username, nil
}

func (u *UserUsecase) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	user, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}
