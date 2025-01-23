package controller

import (
	"net/http"

	auth "auth/domain"
	user1 "user/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	AuthUsecase auth.AuthUsecase
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(authUsecase auth.AuthUsecase) *AuthController {
	return &AuthController{AuthUsecase: authUsecase}
}

// SignUp handles user registration
// @Summary Register a new user
// @Description Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body user1.RegisterUser true "User registration data"
// @Success 200 {object} user1.ResponseUser
// @Failure 400 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/signup [post]
func (ac *AuthController) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input user1.RegisterUser
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Initialize validator
		validate := validator.New()

		// Validate the input struct
		if err := validate.Struct(input); err != nil {
			// Collect and return validation errors
			var errors []string
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, err.Field()+" failed validation: "+err.Tag())
			}

			c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}

		// Register the user
		user, err := ac.AuthUsecase.RegisterUser(input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user":    user1.CreateResponseUser(user),
			"message": "User registered successfully",
		})
	}
}

// LogIn handles user login
// @Summary Log in a user
// @Description Log in a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body user1.LogINUser true "User login data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} httputil.HTTPError
// @Failure 401 {object} httputil.HTTPError
// @Router /api/login [post]
func (ac *AuthController) LogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input user1.LogINUser
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Authenticate the user
		user, accessToken, refreshToken, err := ac.AuthUsecase.LoginUser(input.Email, input.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Check if the user is verified
		if !user.IsVerified {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "you need to be verified", "user": user1.CreateResponseUser(user)})
			return
		}

		// Set the refresh token as a cookie
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "refresh_token",
			Value:    refreshToken,
			Path:     "/",
			HttpOnly: true,
		})

		c.JSON(http.StatusOK, gin.H{
			"user":         user,
			"access_token": accessToken,
		})
	}
}

// LogOut handles user logout
// @Summary Log out a user
// @Description Log out a user
// @Tags Auth
// @Success 200 {string} string "Logged out successfully"
// @Router /api/logout [post]
func (ac *AuthController) LogOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Clear the refresh token cookie
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "refresh_token",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
		})

		c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
	}
}

// Refresh handles token refresh
// @Summary Refresh access token
// @Description Refresh access token using refresh token
// @Tags Auth
// @Success 200 {object} map[string]string
// @Failure 401 {object} httputil.HTTPError
// @Router /api/refresh [post]
func (ac *AuthController) Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the refresh token from the cookie
		cookie, err := c.Request.Cookie("refresh_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No refresh token provided"})
			return
		}

		refreshToken := cookie.Value

		// Refresh the tokens
		accessToken, newRefreshToken, err := ac.AuthUsecase.RefreshTokens(refreshToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Set the new refresh token as a cookie
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "refresh_token",
			Value:    newRefreshToken,
			Path:     "/",
			HttpOnly: true,
		})

		c.JSON(http.StatusOK, gin.H{
			"access_token": accessToken,
		})
	}
}
