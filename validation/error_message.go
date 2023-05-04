package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GetErrorMsg(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required",
			err.Field())
	case "email":
		return fmt.Sprintf("%s is not valid email",
			err.Field())
	case "min":
		return fmt.Sprintf("%s minimum length must be %s",
			err.Field(), err.Param())
	case "gte":
		return fmt.Sprintf("%s value must be greater than equal to %s",
			err.Field(), err.Param())
	}
	return "Unknown error"
}
