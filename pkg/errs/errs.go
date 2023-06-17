package errs

import (
	"fmt"
	"time"

	validator "github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

const (
	ErrCodeUnauthorized        = "UNAUTHORIZED"
	ErrCodeInvalidRequest      = "INVALID_REQUEST"
	ErrCodeInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrCodeForbidden           = "FORBIDDEN"
)

type Error struct {
	Code      string                 `json:"code"`
	Message   string                 `json:"message"`
	Info      map[string]interface{} `json:"info,omitempty"`
	Timestamp Timestamp              `json:"timestamp"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("%s: %s", err.Code, err.Message)
}

func ErrorResponse(err error, code string) Error {
	if errCast, ok := err.(*Error); ok {
		return *errCast
	}
	return Error{
		Code:      code,
		Message:   err.Error(),
		Timestamp: Timestamp(time.Now().UTC()),
	}
}

func ErrorStructResponse(err error) Error {
	return Error{
		Code:      ErrCodeInvalidRequest,
		Message:   "Invalid request",
		Info:      ErrorStruct(err),
		Timestamp: Timestamp(time.Now().UTC()),
	}
}

func ErrorStruct(err error) map[string]interface{} {
	result := make(map[string]interface{})
	if errCast, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errCast {
			result[strcase.ToLowerCamel(e.Field())] = ValidationErrorToText(e)
		}

		return result
	}

	result["NonFieldsError"] = err.Error()
	return result
}

func ValidationErrorToText(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", strcase.ToLowerCamel(e.Field()))
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", strcase.ToLowerCamel(e.Field()), e.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", strcase.ToLowerCamel(e.Field()), e.Param())
	case "email":
		return fmt.Sprintf("Invalid email format")
	case "len":
		return fmt.Sprintf("%s must be %s characters long", strcase.ToLowerCamel(e.Field()), e.Param())
	case "oneof":
		return fmt.Sprintf("%s must be %s", strcase.ToLowerCamel(e.Field()), e.Param())
	}

	return fmt.Sprintf("%s is not valid", strcase.ToLowerCamel(e.Field()))
}
