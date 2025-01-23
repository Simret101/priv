package domain

import (
	"user/domain"
)

type PasswordService interface {
	HashPassword(string) (string, error)
	ComparePassword(string, string) (bool, error)
}

type TokenService interface {
	GenerateAccessToken(user domain.User) (string, error)
	GenerateRefreshToken(user domain.User) (string, error)
	ValidateAccessToken(token string) (*domain.User, error)
	ValidateRefreshToken(token string) (*domain.User, error)
}
