package routes

import (
	"log"
	"os"

	"auth/middleware"
	tokenservice "auth/token_service"
	"user/controller"
	"user/repository"
	"user/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func NewUploadRoute(group *gin.RouterGroup, user_repo repository.UserRepository, rabbitMQConn *amqp.Connection) {
	// Correct the call to NewUploadRepository by passing the required dependencies
	repo := repository.NewUploadRepository(user_repo)
	uc := usecase.NewUploadUsecase(*repo)       // Pass the dereferenced repo
	ctrl := controller.NewUploadController(*uc) // Pass the dereferenced usecase

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Panic(err.Error())
	}
	access_secret := os.Getenv("ACCESSTOKENSECRET")
	if access_secret == "" {
		log.Panic("No access token secret")
	}

	refresh_secret := os.Getenv("REFRESHTOKENSECRET")
	if refresh_secret == "" {
		log.Panic("No refresh token secret")
	}

	// Initialize the token service
	TokenSvc := tokenservice.NewTokenService(access_secret, refresh_secret)

	// Create LoggedIn middleware for authentication
	LoggedInmiddleWare := middleware.LoggedIn(*TokenSvc)

	// Define the route for uploading an image
	group.POST("api/upload/:id", LoggedInmiddleWare, ctrl.UploadImg())
}
