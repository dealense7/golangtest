package requests

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type ValidationRequest interface {
	Validate() []ApiError
}

var validationMessages = map[string]string{
	"required":    "is required",
	"min":         "should be greater than %s",
	"max":         "should be less than %s",
	"len":         "should be %s characters long",
	"eq":          "should be equal to %s",
	"ne":          "should not be equal to %s",
	"lt":          "should be less than %s",
	"lte":         "should be less than or equal to %s",
	"gt":          "should be greater than %s",
	"gte":         "should be greater than or equal to %s",
	"email":       "should be a valid email address",
	"alpha":       "should contain only alphabetic characters",
	"alphanum":    "should contain only alphanumeric characters",
	"numeric":     "should be a valid numeric value",
	"hexadecimal": "should be a valid hexadecimal value",
	"hexcolor":    "should be a valid hexadecimal color code",
	"rgb":         "should be a valid RGB color code",
	"rgba":        "should be a valid RGBA color code",
	"url":         "should be a valid URL",
}

type ApiError struct {
	Field string
	Msg   string
}

func GenerateValidationErrors(request interface{}, err error) []ApiError {

	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ApiError, len(ve))

			for i, fe := range ve {
				field, _ := reflect.TypeOf(request).FieldByName(fe.Field())
				out[i] = ApiError{
					Field: field.Tag.Get("json"),
					Msg:   validationMessages[fe.Tag()],
				}
			}

			return out
		}
		return nil
	}
	return nil
}
