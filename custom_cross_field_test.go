package Golang_Validation

import (
	"github.com/go-playground/validator/v10"
	"strings"
	"testing"
)

func MustIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("MustIgnoreCase: GetStructFieldOK2 failed")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestMustIgnoreCase(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("must_ignore_case", MustIgnoreCase)
	if err != nil {
		panic(err)
	}
	type User struct {
		Username string `validate:"required,must_ignore_case=Email|must_ignore_case=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
		Name     string `validate:"required"`
	}
	user := User{
		Username: "",
		Email:    "",
		Phone:    "",
		Name:     "",
	}
	err = validate.Struct(user)
	if err != nil {
		t.Error(err)
	}
}
