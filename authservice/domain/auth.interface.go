package domain

import (
	"user/domain"

	"github.com/gin-gonic/gin"
)

// AuthController defines the methods that an authentication controller should implement
//
//go:generate mockgen
type AuthController interface {
	SignUp() gin.HandlerFunc
	LogIn() gin.HandlerFunc
	GoogleLogIn() gin.HandlerFunc
	GoogleCallBack() gin.HandlerFunc
	LogOut() gin.HandlerFunc
	Refresh() gin.HandlerFunc
}

// AuthUsecase defines the methods that an authentication use case should implement
type AuthUsecase interface {
	RegisterUser(domain.RegisterUser) (domain.User, error)
	LoginUser(string, string) (domain.User, string, string, error)

	RefreshTokens(string) (string, string, error)
}

// AuthRepository defines the methods that an authentication repository should implement
type AuthRepository interface {
	SaveUser(*domain.User) error
	FindUserByEmail(string) (*domain.User, error)
}


