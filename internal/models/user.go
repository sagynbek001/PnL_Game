package models

import (
	"context"
)

// Users ...

type User struct {
	ID           uint64 `json:"id"`
	Login        string `json:"login"`
	PasswordHash string `json:"password_hash"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// UsersRepo represents Users repository contract
type UserUsecase interface {
	Create(ctx context.Context, user *User) error
	GetAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id uint64) (User, error)
	Update(ctx context.Context, users *User) error
	Delete(ctx context.Context, id uint64) error
	FindByLogin(ctx context.Context, email string) (User, error)
}

// UsersRepo represents Users repository contract
type UserRepo interface {
	Create(ctx context.Context, users *User) error
	GetAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id uint64) (User, error)
	Update(ctx context.Context, users *User) error
	Delete(ctx context.Context, id uint64) error
	FindByLogin(ctx context.Context, email string) (User, error)
}
