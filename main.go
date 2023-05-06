package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhm-rai/mygram/config"
	"github.com/ilhm-rai/mygram/controller"
	"github.com/ilhm-rai/mygram/exception"
	"github.com/ilhm-rai/mygram/repository"
	"github.com/ilhm-rai/mygram/service"

	_ "github.com/ilhm-rai/mygram/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyGram API
// @version 1.0
// @description MyGram is a simple API for Final Project DTS Kominfo
// @termOfService http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	configuration := config.New()
	database := config.NewPostgresDatabase(configuration)

	userRepository := repository.NewUserRepository(database)
	photoRepository := repository.NewPhotoRepository(database)
	socialMediaRepository := repository.NewSocialMediaRepository(database)
	commentRepository := repository.NewCommentRepository(database)

	authService := service.NewAuthService(&userRepository)
	photoService := service.NewPhotoService(&photoRepository)
	socialMediaService := service.NewSocialMediaService(&socialMediaRepository)
	commentService := service.NewCommentService(&commentRepository)

	authController := controller.NewAuthController(&authService)
	photoController := controller.NewPhotoController(&photoService)
	socialMediaController := controller.NewSocialMediaController(&socialMediaService)
	commentController := controller.NewCommentController(&commentService)

	app := gin.Default()

	v1 := app.Group("api/v1")
	{
		authController.Route(v1)
		photoController.Route(v1)
		socialMediaController.Route(v1)
		commentController.Route(v1)
	}

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err := app.Run(":" + configuration.Get("PORT"))
	exception.PanicIfNeeded(err)
}
