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

type PhotoController struct {
	PhotoService service.PhotoService
}

func NewPhotoController(photoService *service.PhotoService) PhotoController {
	return PhotoController{
		PhotoService: *photoService,
	}
}

func (controller *PhotoController) Route(app *gin.Engine) {
	photoRouter := app.Group("photos")
	{
		photoRouter.Use(authentication())
		photoRouter.GET("/", controller.FindPhotos)
		photoRouter.POST("/", controller.CreatePhoto)
		photoRouter.GET("/:id", controller.FindPhoto)
		photoRouter.PUT("/:id", middleware.PhotoAuthorization(controller.PhotoService), controller.UpdatePhoto)
		photoRouter.DELETE("/:id", middleware.PhotoAuthorization(controller.PhotoService), controller.DeletePhoto)
	}
}

func (controller *PhotoController) FindPhotos(c *gin.Context) {
	photos, err := controller.PhotoService.FindAll()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Data:   photos,
		Errors: nil,
	})
}

func (controller *PhotoController) FindPhoto(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid parameter",
		})
		return
	}

	photo, err := controller.PhotoService.FindById(id)

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

func (controller PhotoController) CreatePhoto(c *gin.Context) {
	var request model.SavePhotoRequest

	if !validateRequest(c, &request) {
		return
	}

	userData := c.MustGet("UserData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	request.UserId = userId

	err := controller.PhotoService.Save(request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Data:   "Photo successfully created",
		Errors: nil,
	})
}

func (controller PhotoController) UpdatePhoto(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var request model.SavePhotoRequest

	if !validateRequest(c, &request) {
		return
	}

	userData := c.MustGet("UserData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	request.ID = uint(id)
	request.UserId = userId
	err := controller.PhotoService.Save(request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Data:   "Photo successfully updated",
		Errors: nil,
	})
}

func (controller PhotoController) DeletePhoto(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := controller.PhotoService.DeleteById(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Data:   "Photo successfully deleted",
		Errors: nil,
	})
}
