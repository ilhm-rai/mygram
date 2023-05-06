package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhm-rai/mygram/model"
	"github.com/ilhm-rai/mygram/service"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService *service.AuthService) AuthController {
	return AuthController{
		AuthService: *authService,
	}
}

func (controller *AuthController) Route(app *gin.RouterGroup) {
	app.POST("/login", controller.Login)
	app.POST("/register", controller.Register)
}

// Register godoc
// @Summary Register new user
// @Description MyGram new user registration
// @Tags auth
// @Accept json
// @Produce json
// @Param model.RegisterUserRequest body model.RegisterUserRequest true "register request"
// @Success 200 {object} model.WebResponse
// @Failure 500 {object} model.ErrResponse
// @Router /register [post]
func (controller *AuthController) Register(c *gin.Context) {
	var request model.RegisterUserRequest

	valid := validateRequest(c, &request)

	if !valid {
		return
	}

	if err := controller.AuthService.Register(request); err != nil {
		var Errors []model.ErrorMsg
		if err.Error() == "email_or_username_exist" {
			Errors = append(Errors, model.ErrorMsg{Field: "Email or username", Message: "Email or username already registered"})
			c.AbortWithStatusJSON(http.StatusBadRequest, model.WebResponse{
				Code:   http.StatusBadRequest,
				Data:   nil,
				Errors: Errors,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Data:   "Successfully registered",
		Errors: nil,
	})
}

// Login godoc
// @Summary Login
// @Description Login to MyGram
// @Tags auth
// @Accept json
// @Produce json
// @Param model.LoginUserRequest body model.LoginUserRequest true "login request"
// @Success 200 {object} model.WebResponse
// @Failure 500 {object} model.ErrResponse
// @Router /login [post]
func (controller *AuthController) Login(c *gin.Context) {
	var request model.LoginUserRequest

	valid := validateRequest(c, &request)
	if !valid {
		return
	}

	token, err := controller.AuthService.Login(request)

	if err != nil {
		if err.Error() == "user_not_found" {
			c.AbortWithStatusJSON(http.StatusNotFound, model.ErrResponse{
				Code:    http.StatusNotFound,
				Message: "Email or username not found",
			})
			return
		}
		if err.Error() == "invalid_password" {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrResponse{
				Code:    http.StatusBadRequest,
				Message: "Password is invalid",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code: http.StatusOK,
		Data: gin.H{
			"token": token,
		},
		Errors: nil,
	})
}
