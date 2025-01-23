package domain

import (
	"user/domain"

	"github.com/gin-gonic/gin"
)

// VerifyEmail_Repository_interface defines the methods that an email verification repository should implement
type VerifyEmail_Repository_interface interface {
	VerifyUser(id string) error
}

// VerifyEmail_Usecase_interface defines the methods that an email verification use case should implement
type VerifyEmail_Usecase_interface interface {
	SendVerifyEmail(id string, vuser domain.VerifyEmail) error
	VerifyUser(token string) error
	SendForgretPasswordEmail(id string, vuser domain.VerifyEmail) error
	ValidateForgetPassword(id string, token string) error
}

// VerifyEmail_Controller_interface defines the methods that an email verification controller should implement
type VerifyEmail_Controller_interface interface {
	SendVerificationEmail() gin.HandlerFunc
	VerifyEmail() gin.HandlerFunc
}
