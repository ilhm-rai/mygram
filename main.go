package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhm-rai/mygram/config"
	"github.com/ilhm-rai/mygram/controller"
	"github.com/ilhm-rai/mygram/exception"
	"github.com/ilhm-rai/mygram/repository"
	"github.com/ilhm-rai/mygram/service"
)

func main() {
	configuration := config.New()
	database := config.NewPostgresDatabase(configuration)

	userRepository := repository.NewUserRepository(database)
	photoRepository := repository.NewPhotoRepository(database)
	socialMediaRepository := repository.NewSocialMediaRepository(database)

	authService := service.NewAuthService(&userRepository)
	photoService := service.NewPhotoService(&photoRepository)
	socialMediaService := service.NewSocialMediaService(&socialMediaRepository)

	authController := controller.NewAuthController(&authService)
	photoController := controller.NewPhotoController(&photoService)
	socialMediaController := controller.NewSocialMediaController(&socialMediaService)

	app := gin.Default()

	authController.Route(app)
	photoController.Route(app)
	socialMediaController.Route(app)

	err := app.Run(":" + configuration.Get("PORT"))
	exception.PanicIfNeeded(err)
}
