package controller

import (
	"fmt"
	"net/http"

	auth "auth/domain"
	user "user/domain"

	"github.com/gin-gonic/gin"
)

// EmailVControler handles email verification related operations
type EmailVControler struct {
	user_usecase user.User_Usecase_interface
	email_uc     auth.VerifyEmail_Usecase_interface
}

// NewEmailVController creates a new instance of EmailVControler
func NewEmailVController(email_usecase auth.VerifyEmail_Usecase_interface, user_usecase user.User_Usecase_interface) *EmailVControler {
	return &EmailVControler{
		email_uc:     email_usecase,
		user_usecase: user_usecase,
	}
}

// SendVerificationEmail sends a verification email to the user
// @Summary Send a verification email
// @Description Send a verification email to the user
// @Tags Email Verification
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param email body user.VerifyEmail true "Email data"
// @Success 202 {string} string "email sent to: {email}"
// @Failure 400 {object} httputil.HTTPError
// @Router /api/verify-email/{id} [post]
func (ctrl *EmailVControler) SendVerificationEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var model user.VerifyEmail
		if err := ctx.BindJSON(&model); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id := ctx.Param("id")
		if err := ctrl.email_uc.SendVerifyEmail(id, model); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"errorss": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusAccepted, gin.H{"message": fmt.Sprintf("email sent to: %s", model.Email)})
	}
}

// VerifyEmail verifies the user's email using the provided token
// @Summary Verify email
// @Description Verify the user's email using the provided token
// @Tags Email Verification
// @Accept json
// @Produce json
// @Param token path string true "Verification Token"
// @Success 202 {string} string "Verified"
// @Failure 400 {object} httputil.HTTPError
// @Router /api/verify-email/{token} [get]
func (ctrl *EmailVControler) VerifyEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Param("token")
		err := ctrl.email_uc.VerifyUser(token)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusAccepted, gin.H{"message": "Verified"})
	}
}

// ForgetPasswordValidate validates the forget password token and resets the password
// @Summary Validate forget password token
// @Description Validate the forget password token and reset the password
// @Tags Password Reset
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Param token query string true "Token"
// @Success 202 {object} user.User
// @Failure 400 {object} httputil.HTTPError
// @Router /api/forget-password/ [get]
func (ctrl *EmailVControler) ForgetPasswordValidate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Query("id")
		token := ctx.Query("token")

		err := ctrl.email_uc.ValidateForgetPassword(id, token)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		update_password := user.UpdatePassword{
			Password:        "12345678",
			ConfirmPassword: "12345678",
		}
		user, err := ctrl.user_usecase.UpdatePassword(id, update_password)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusAccepted, gin.H{"user": user, "message": "your password is reset to 12345678, you can change it anytime you want"})
	}
}

// SendForgetPasswordEmail sends a forget password email to the user
// @Summary Send forget password email
// @Description Send a forget password email to the user
// @Tags Password Reset
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param email body user.VerifyEmail true "Email data"
// @Success 202 {string} string "email sent to: {email}"
// @Failure 400 {object} httputil.HTTPError
// @Router /api/forget-password/{id} [post]
func (ctrl *EmailVControler) SendForgetPasswordEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var model user.VerifyEmail
		if err := ctx.BindJSON(&model); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id := ctx.Param("id")
		if err := ctrl.email_uc.SendForgretPasswordEmail(id, model); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"errorss": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusAccepted, gin.H{"message": fmt.Sprintf("email sent to: %s", model.Email)})
	}
}
