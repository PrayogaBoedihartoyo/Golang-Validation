package Golang_Validation

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestOrRule(t *testing.T) {
	type Login struct {
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required"`
	}
	login := Login{
		Username: "",
		Password: "",
	}
	validate := validator.New()
	err := validate.Struct(login)
	if err != nil {
		t.Error(err.Error())
	}
}
