package validations

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Struct(body interface{}) *[]string {
	if err := validate.Struct(body); err != nil {
		validationErrors := strings.Split(err.Error(), "\n")

		return &validationErrors
	}
	return nil
}
