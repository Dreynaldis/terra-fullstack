package repository

import (
	"backend/internal/model"
	"context"
	"database/sql"
	"fmt"

	"errors"
)

type UserRepository struct {
	queries *model.Queries
}

func NewUserRepository(queries *model.Queries) *UserRepository {
	return &UserRepository{
		queries: queries,
	}
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *UserRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return user != (model.User{}), nil
}
func (r *UserRepository) UsernameExists(ctx context.Context, username string) (bool, error) {
	user, err := r.queries.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return user != (model.User{}), nil
}

func (r *UserRepository) CreateUser(ctx context.Context, username, email, password string) error {
	usernameExists, err := r.UsernameExists(ctx, username)
	if err != nil {
		return fmt.Errorf("error checking if username exists: %w", err)
	}
	if usernameExists {
		return errors.New("username already exists")
	}
	emailExists, err := r.EmailExists(ctx, email)
	if err != nil {
		return fmt.Errorf("error checking if email exists: %w", err)
	}
	if emailExists {
		return errors.New("email already exists")
	}
	err = r.queries.CreateUser(ctx, model.CreateUserParams{
		Username: username,
		Email:    email,
		Password: password,
	})
	if err != nil {
		fmt.Printf("error creating user:%v\n ", err)
		return err
	}

	return nil
}
