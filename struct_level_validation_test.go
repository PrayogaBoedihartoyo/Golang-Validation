package Golang_Validation

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

type RegisterRequest struct {
	Username        string `validate:"required"`
	Email           string `validate:"required,email"`
	Password        string `validate:"required,min=6"`
	ConfirmPassword string `validate:"required,eqfield=Password"`
}

func MustValidRegisterSuccess(level validator.StructLevel) {
	registerRequest := level.Current().Interface().(RegisterRequest)
	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Password {
		level.ReportError(registerRequest.Username, "Username", "Username", "username", "")
	} else {
		level.ReportError(registerRequest.Email, "Email", "Email", "email", "")
	}
}

func TestMustValidRegisterSuccess(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(MustValidRegisterSuccess, RegisterRequest{})

	registerRequest := RegisterRequest{
		Username:        "tag",
		Email:           "yoga@example.com",
		Password:        "secret",
		ConfirmPassword: "secret",
	}
	err := validate.Struct(registerRequest)
	if err != nil {
		t.Error(err)
	}
}
