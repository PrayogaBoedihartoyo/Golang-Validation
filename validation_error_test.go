package Golang_Validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestValidationError(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=6,max=12"`
	}
	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "Prayoga@gmail.com",
		Password: "Boedihartoyo",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		validateError := err.(validator.ValidationErrors)
		for _, fieldError := range validateError {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.ActualTag())
		}
	}
}
