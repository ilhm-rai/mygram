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

type CommentController struct {
	CommentService service.CommentService
}

func NewCommentController(photoService *service.CommentService) CommentController {
	return CommentController{
		CommentService: *photoService,
	}
}

func (controller *CommentController) Route(app *gin.RouterGroup) {
	photoRouter := app.Group("comments")
	{
		photoRouter.Use(authentication())
		photoRouter.GET("/", controller.FindComments)
		photoRouter.POST("/", controller.CreateComment)
		photoRouter.GET("/:id", controller.FindComment)
		photoRouter.PUT("/:id", middleware.CommentAuthorization(controller.CommentService), controller.UpdateComment)
		photoRouter.DELETE("/:id", middleware.CommentAuthorization(controller.CommentService), controller.DeleteComment)
	}
}

// FindComments godoc
// @Summary Find all comments
// @Description Find all comments of all photos
// @Tags comments
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Success 200 {object} model.WebResponse
// @Failure 500 {object} model.ErrResponse
// @Router /comments [get]
func (controller *CommentController) FindComments(c *gin.Context) {
	comments, err := controller.CommentService.FindAll()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Data:   comments,
		Errors: nil,
	})
}

// FindComment godoc
// @Summary Find comment by id
// @Description Find a comment identified by the given id
// @Tags comments
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param id path int true "ID of the comment"
// @Success 200 {object} model.WebResponse
// @Failure 404 {object} model.ErrResponse
// @Failure 500 {object} model.ErrResponse
// @Router /comments/{id} [get]
func (controller *CommentController) FindComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid parameter",
		})
		return
	}

	photo, err := controller.CommentService.FindById(id)

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

// CreateComment godoc
// @Summary Create a new comment
// @Description Create a new comment for specific photo
// @Tags comments
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param model.SaveCommentRequest body model.SaveCommentRequest true "create comment request"
// @Success 200 {object} model.WebResponse
// @Failure 500 {object} model.ErrResponse
// @Router /comments [post]
func (controller CommentController) CreateComment(c *gin.Context) {
	var request model.SaveCommentRequest

	if !validateRequest(c, &request) {
		return
	}

	userData := c.MustGet("UserData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	request.UserId = userId

	err := controller.CommentService.Save(request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Data:   "Comment successfully created",
		Errors: nil,
	})
}

// UpdateComment godoc
// @Summary Update a comment
// @Description Update a comment identified by the given id
// @Tags comments
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param id path int true "ID of the comment to be updated"
// @Param model.SaveCommentRequest body model.SaveCommentRequest true "update comment request"
// @Success 200 {object} model.WebResponse
// @Failure 404 {object} model.ErrResponse
// @Failure 500 {object} model.ErrResponse
// @Router /comments/{id} [put]
func (controller CommentController) UpdateComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var request model.SaveCommentRequest

	if !validateRequest(c, &request) {
		return
	}

	userData := c.MustGet("UserData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	request.ID = uint(id)
	request.UserId = userId
	err := controller.CommentService.Save(request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Data:   "Comment successfully updated",
		Errors: nil,
	})
}

// DeleteComment godoc
// @Summary Delete a comment
// @Description Delete a comment identified by the given id
// @Tags comments
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param id path int true "ID of the comment to be deleted"
// @Success 200 {object} model.WebResponse
// @Failure 404 {object} model.ErrResponse
// @Failure 500 {object} model.ErrResponse
// @Router /comments/{id} [delete]
func (controller CommentController) DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := controller.CommentService.DeleteById(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Data:   "Comment successfully deleted",
		Errors: nil,
	})
}
