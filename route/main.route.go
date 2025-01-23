package route

import (
	captcha "auth/captha" // Import the CAPTCHA package
	auth "auth/config"
	auth1 "auth/database"
	auth3 "auth/route"
	blog1 "blog/database"
	blog4 "blog/route"
	user1 "user/database"
	user4 "user/route"
	user5 "user/repository"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp" // Import RabbitMQ package
)

func SetUp(router *gin.Engine) {
	var clinect auth.ServerConnection
	clinect.Connect_could()

	userCollection := &user1.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("Users"),
	}

	authCollection := &auth1.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("Users"),
	}
	blogCollection := &blog1.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("Blogs"),
	}

	stateCollection := &auth1.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("States"),
	}
	bookmarkCollection := &blog1.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("Bookmarks"),
	}

	// Initialize RabbitMQ connection
	rabbitMQConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic("Failed to connect to RabbitMQ: " + err.Error())
	}
	defer rabbitMQConn.Close()

	userRoute := router.Group("")
	verifiRoute := router.Group("")
	authRoute := router.Group("")
	blogRot := router.Group("")
	uploadRoute := router.Group("")
	bookmarkRoute := router.Group("")

	userRepo := user5.NewUserRepository(userCollection)

	// Add CAPTCHA routes
	captchaRoute := router.Group("/captcha")
	captchaRoute.GET("/", func(c *gin.Context) {
		captcha.GenerateCaptcha(c)
	})

	auth3.NewVerifyEmialRoute(verifiRoute, userCollection)
	blog4.NewBlogRoutes(blogRot, blogCollection, userCollection)
	user4.NewUserRoute(userRoute, userCollection)
	auth3.NewAuthRoute(authRoute, authCollection, stateCollection)

	// Pass RabbitMQ connection to the upload route
	user4.NewUploadRoute(uploadRoute, *userRepo, rabbitMQConn)

	blog4.NewBookmarkRoutes(bookmarkRoute, bookmarkCollection, blogCollection)
}
