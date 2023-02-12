package Golang_Validation

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestDoubleValidation(t *testing.T) {
	validate := validator.New()
	password := "123456"
	confirmPassword := "123456"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		t.Error(err)
	}
}
