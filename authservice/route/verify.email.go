package routes

import (
	"log"
	"os"

	"auth/controller"
	auth "auth/repository"
	auth3 "auth/usecase"
	user1 "user/database"
	user "user/repository"
	user3 "user/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func NewVerifyEmialRoute(group *gin.RouterGroup, user_collection user1.CollectionInterface) {
	repo := user.NewUserRepository(user_collection)
	user_usecase := user3.NewUserUseCase(repo)

	email_repo := auth.NewEmailVRepo(*repo)
	email_usecase := auth3.NewEmailVUsecase(user_usecase, email_repo)
	email_ctrl := controller.NewEmailVController(email_usecase, user_usecase)

	//load middlewares
	err := godotenv.Load()
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
	//TokenSvc := *tokenservice.NewTokenService(access_secret, refresh_secret)

	//LoggedInmiddleWare := middleware.LoggedIn(TokenSvc)

	group.POST("api/verify-email/:id", email_ctrl.SendVerificationEmail())
	group.GET("api/verify-email/:token", email_ctrl.VerifyEmail())

	group.POST("api/forget-password/:id", email_ctrl.SendForgetPasswordEmail())
	group.GET("/api/forget-password/", email_ctrl.ForgetPasswordValidate())
}
