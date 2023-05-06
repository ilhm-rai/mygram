package controller

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ilhm-rai/mygram/middleware"
	"github.com/ilhm-rai/mygram/model"
	"github.com/ilhm-rai/mygram/service"
)

type SocialMediaController struct {
	SocialMediaService service.SocialMediaService
}

func NewSocialMediaController(photoService *service.SocialMediaService) SocialMediaController {
	return SocialMediaController{
		SocialMediaService: *photoService,
	}
}

func (controller *SocialMediaController) Route(app *gin.Engine) {
	photoRouter := app.Group("social-media")
	{
		photoRouter.Use(authentication())
		photoRouter.GET("/", controller.FindSocialMedia)
		photoRouter.POST("/", controller.CreateSocialMedia)
		photoRouter.GET("/:id", controller.FindSocialMediaById)
		photoRouter.PUT("/:id", middleware.SocialMediaAuthorization(controller.SocialMediaService), controller.UpdateSocialMedia)
		photoRouter.DELETE("/:id", middleware.SocialMediaAuthorization(controller.SocialMediaService), controller.DeleteSocialMedia)
	}
}

func (controller *SocialMediaController) FindSocialMedia(c *gin.Context) {
	socialMedia, err := controller.SocialMediaService.FindAll()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Data:   socialMedia,
		Errors: nil,
	})
}

func (controller *SocialMediaController) FindSocialMediaById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid parameter",
		})
		return
	}

	photo, err := controller.SocialMediaService.FindById(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Data:   photo,
		Errors: nil,
	})
}

func (controller SocialMediaController) CreateSocialMedia(c *gin.Context) {
	var request model.SaveSocialMediaRequest

	if !validateRequest(c, &request) {
		return
	}

	userData := c.MustGet("UserData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	request.UserId = userId

	err := controller.SocialMediaService.Save(request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Data:   "Social media successfully created",
		Errors: nil,
	})
}

func (controller SocialMediaController) UpdateSocialMedia(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var request model.SaveSocialMediaRequest

	if !validateRequest(c, &request) {
		return
	}

	userData := c.MustGet("UserData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	request.ID = uint(id)
	request.UserId = userId
	err := controller.SocialMediaService.Save(request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Data:   "Social media successfully updated",
		Errors: nil,
	})
}

func (controller SocialMediaController) DeleteSocialMedia(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := controller.SocialMediaService.DeleteById(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Data:   "Social media successfully deleted",
		Errors: nil,
	})
}
