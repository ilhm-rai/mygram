package helper

import (
	"errors"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	key = os.Getenv("SECRET_KEY")
)

func GenerateToken(id uint, email string, username string) string {
	claims := jwt.MapClaims{
		"id":       id,
		"email":    email,
		"username": username,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(key))

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	err := errors.New("not_authorized")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, err
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(key), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}
