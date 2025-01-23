package routes

import (
	"log"
	"os"

	"auth/domain"
	"auth/middleware"
	tokenservice "auth/token_service"
	"user/controller"
	"user/database"
	"user/repository"
	"user/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func NewUserRoute(group *gin.RouterGroup, user_collection database.CollectionInterface) {
	// Initialize repository, use case, and controller
	repo := repository.NewUserRepository(user_collection)
	usecase := usecase.NewUserUseCase(repo)
	ctrl := controller.NewUserController(usecase)

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Panic(err.Error())
	}
	accessSecret := os.Getenv("ACCESSTOKENSECRET")
	if accessSecret == "" {
		log.Panic("ACCESSTOKENSECRET is missing")
	}

	refreshSecret := os.Getenv("REFRESHTOKENSECRET")
	if refreshSecret == "" {
		log.Panic("REFRESHTOKENSECRET is missing")
	}

	// Initialize Token Service
	tokenSvc := tokenservice.NewTokenService(accessSecret, refreshSecret)

	// Define middlewares
	loggedInMiddleware := middleware.LoggedIn(*tokenSvc)
	adminMiddleware := middleware.RoleBasedAccess(accessSecret, []domain.Role{domain.Role("admin")}) // Only Admin has access

	// Define user routes
	group.GET("api/user/:id", loggedInMiddleware, ctrl.GetOneUser())
	group.GET("api/user/", loggedInMiddleware, ctrl.GetUsers())

	group.PUT("api/user/:id", loggedInMiddleware, ctrl.UpdateUser())
	group.DELETE("api/user/:id", loggedInMiddleware, ctrl.DeleteUser())

	// Routes restricted to Admin role
	group.PUT("api/demote/:id", loggedInMiddleware, adminMiddleware, ctrl.DemoteUser())
	group.PUT("api/promote/:id", loggedInMiddleware, adminMiddleware, ctrl.PromoteUser())
}
