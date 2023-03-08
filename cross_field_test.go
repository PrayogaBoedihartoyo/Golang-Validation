package Golang_Validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestCrossField(t *testing.T) {
	type RegisterUser struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=6,max=12"`
		ConfirmPassword string `validate:"required,eqfield=Password"`
	}
	validate := validator.New()
	registerUser := RegisterUser{
		Username:        "Prayoga",
		Password:        "123",
		ConfirmPassword: "123",
	}
	err := validate.Struct(registerUser)
	if err != nil {
		fmt.Println(err.Error())
	}
}
