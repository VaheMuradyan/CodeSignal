package utils

import (
	"errors"
	"time"

	"codesignal.com/example/gin/todoapp/models"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notpast", notPastDate)
		v.RegisterValidation("maxlength", maxLengthTitle)
	}
}

var notPastDate validator.Func = func(f1 validator.FieldLevel) bool {
	data, ok := f1.Field().Interface().(time.Time)
	if !ok {
		return false
	}

	return !data.Before(time.Now().Add(-1 * time.Second))
}

var maxLengthTitle validator.Func = func(f1 validator.FieldLevel) bool {
	word, ok := f1.Field().Interface().(string)
	if !ok {
		return false
	}

	return len(word) < 50
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

func CheckForDuplicates(todos []models.Todo) []string {
	titleMap := make(map[string]bool)
	dublicates := []string{}

	for _, todo := range todos {
		if _, found := titleMap[todo.Title]; found {
			dublicates = append(dublicates, todo.Title)
		} else {
			titleMap[todo.Title] = true
		}
	}

	return dublicates
}
