package utitlity

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func PhoneNumberValidation(fl validator.FieldLevel) bool {
	fieldVal := fl.Field().String()
	match, _ := regexp.MatchString("^[0-9+-]+$", fieldVal)
	return match
}

func EmailValidation(fl validator.FieldLevel) bool {
	fieldVal := fl.Field().String()
	match, _ := regexp.MatchString("^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$", fieldVal)
	return match
}

func AlphaSpace(fl validator.FieldLevel) bool {
	fieldVal := fl.Field().String()
	match, _ := regexp.MatchString("^[a-zA-Z\\s]+$", fieldVal)
	return match
}

func Date(fl validator.FieldLevel) bool {
	fieldVal := fl.Field().String()
	match, _ := regexp.MatchString("^(0[1-9]|[12][0-9]|3[01])/(0[1-9]|1[012])/((19|20)\\d\\d)$", fieldVal)
	return match
}

func Time(fl validator.FieldLevel) bool {
	fieldVal := fl.Field().String()
	match, _ := regexp.MatchString("^([01]?[0-9]|2[0-3]):[0-5][0-9]$", fieldVal)
	return match
}
