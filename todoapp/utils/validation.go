package utils

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notpast", notPastDate)
	}
}

var notPastDate validator.Func = func(f1 validator.FieldLevel) bool {
	data, ok := f1.Field().Interface().(time.Time)
	if !ok {
		return false
	}

	return !data.Before(time.Now().Add(-1 * time.Second))
}

func EnhancedErrorMessages(err error) map[string]string {
	var ve validator.ValidationErrors
	out := make(map[string]string)

	if errors.As(err, &ve) {
		for _, fe := range ve {
			switch fe.Tag() {
			case "required":
				out[fe.Field()] = "This field is required"
			case "notpast":
				out[fe.Field()] = "The date must not be in the past"
			}
		}
	}
	return out
}
