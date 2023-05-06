package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ilhm-rai/mygram/helper"
	"github.com/ilhm-rai/mygram/model"
	"github.com/ilhm-rai/mygram/validation"
)

var (
	validate *validator.Validate
)

func validateRequest(c *gin.Context, request interface{}) bool {
	if validate == nil {
		validate = validator.New()
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return false
	}

	if err := validate.Struct(request); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errs := make([]model.ErrorMsg, len(ve))
			for i, fe := range ve {
				errs[i] = model.ErrorMsg{Field: fe.Field(), Message: validation.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, model.WebResponse{
				Code:   http.StatusBadRequest,
				Data:   nil,
				Errors: errs,
			})
		}
		return false
	}
	return true
}

func authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := helper.VerifyToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrResponse{
				Code:    http.StatusUnauthorized,
				Message: "Sign in to proceed",
			})
			return
		}
		ctx.Set("UserData", token)
		ctx.Next()
	}
}
