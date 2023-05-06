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

func (controller *SocialMediaController) Route(app *gin.RouterGroup) {
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

// FindSocialMedia godoc
// @Summary Find all social media
// @Description Find all social media from all users
// @Tags social-media
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Success 200 {object} model.WebResponse
// @Failure 500 {object} model.ErrResponse
// @Router /social-media [get]
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

// FindSocialMediaById godoc
// @Summary Find social media by id
// @Description Find a social media identified by the given id
// @Tags social-media
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param id path int true "ID of the social media"
// @Success 200 {object} model.WebResponse
// @Failure 404 {object} model.ErrResponse
// @Failure 500 {object} model.ErrResponse
// @Router /social-media/{id} [get]
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

// CreateSocialMedia godoc
// @Summary Create a new social media
// @Description Create a new social media for specific user
// @Tags social-media
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param model.SaveSocialMediaRequest body model.SaveSocialMediaRequest true "create photo request"
// @Success 200 {object} model.WebResponse
// @Failure 500 {object} model.ErrResponse
// @Router /social-media [post]
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

// UpdateSocialMedia godoc
// @Summary Update a social media
// @Description Update a social media identified by the given id
// @Tags social-media
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param id path int true "ID of the social media to be updated"
// @Param model.SaveSocialMediaRequest body model.SaveSocialMediaRequest true "update social media request"
// @Success 200 {object} model.WebResponse
// @Failure 404 {object} model.ErrResponse
// @Failure 500 {object} model.ErrResponse
// @Router /social-media/{id} [put]
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

// FindSocialMedia godoc
// @Summary Find social media by id
// @Description Find a social media identified by the given id
// @Tags social-media
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param id path int true "ID of the social media"
// @Success 200 {object} model.WebResponse
// @Failure 404 {object} model.ErrResponse
// @Failure 500 {object} model.ErrResponse
// @Router /social-media/{id} [delete]
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
