package Golang_Validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestMultipleValidation(t *testing.T) {
	validate := validator.New()
	var user string = "123"

	err := validate.Var(user, "required,number")
	if err != nil {
		fmt.Println(err.Error())
	}
}
