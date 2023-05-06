package middleware

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ilhm-rai/mygram/model"
	"github.com/ilhm-rai/mygram/service"
)

func SocialMediaAuthorization(service service.SocialMediaService) gin.HandlerFunc {
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
		socialMedia, err := service.FindById(id)

		if err != nil {
			if err.Error() == "social_media_not_found" {
				c.AbortWithStatusJSON(http.StatusNotFound, model.ErrResponse{
					Code:    http.StatusNotFound,
					Message: "Social media not found",
				})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		if socialMedia.UserId != userId {
			c.AbortWithStatusJSON(http.StatusForbidden, model.ErrResponse{
				Code:    http.StatusForbidden,
				Message: "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}
