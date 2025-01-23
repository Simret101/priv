package routes

import (
	"log"
	"os"

	"auth/controller"
	"auth/database"
	passwordservice "auth/passwordservice"
	"auth/repository"
	tokenservice "auth/token_service"
	"auth/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func NewAuthRoute(group *gin.RouterGroup, users, state database.CollectionInterface) {
	AuthRepo, err := repository.NewAuthRepo(users)
	if err != nil {
		log.Panic(err.Error())
	}

	err = godotenv.Load()
	if err != nil {
		log.Panic(err.Error())
	}

	access_secret := os.Getenv("ACCESSTOKENSECRET")
	if access_secret == "" {
		log.Panic("No accesstoken")
	}

	refresh_secret := os.Getenv("REFRESHTOKENSECRET")
	if refresh_secret == "" {
		log.Panic("No refreshtoken")
	}

	TokenSvc := tokenservice.NewTokenService(access_secret, refresh_secret)
	PasswordSvc := &passwordservice.PasswordS{}

	AuthUsecase := usecase.NewAuthUsecase(AuthRepo, PasswordSvc, TokenSvc)
	AuthController := controller.NewAuthController(AuthUsecase)

	group.POST("/signup", AuthController.SignUp())
	group.POST("/login", AuthController.LogIn())
	group.POST("/logout", AuthController.LogOut())
	group.POST("/refresh", AuthController.Refresh())

}
