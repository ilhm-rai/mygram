package middleware

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ilhm-rai/mygram/model"
	"github.com/ilhm-rai/mygram/service"
)

func CommentAuthorization(service service.CommentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrResponse{
				Code:    http.StatusBadRequest,
				Message: "Invalid parameter",
			})
			return
		}

		userData := c.MustGet("UserData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		photo, err := service.FindById(id)

		if err != nil {
			if err.Error() == "photo_not_found" {
				c.AbortWithStatusJSON(http.StatusNotFound, model.ErrResponse{
					Code:    http.StatusNotFound,
					Message: "Comment not found",
				})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		if photo.UserId != userId {
			c.AbortWithStatusJSON(http.StatusForbidden, model.ErrResponse{
				Code:    http.StatusForbidden,
				Message: "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}
