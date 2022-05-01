package validator

import (
	"regexp"

	"github.com/go-playground/validator"
)

func RegexTag(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	if field == "" {
		return true
	}
	regex := regexp.MustCompile(`^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`)
	match := regex.MatchString(field)
	return match
}

func New() *validator.Validate {
	customValidator := validator.New()
	err := customValidator.RegisterValidation("slug", RegexTag)
	if err != nil {
		panic(err)
	}

	return customValidator
}
