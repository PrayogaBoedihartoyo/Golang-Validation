package Golang_Validation

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		t.Error("validator is nil")
	}
}

func TestValidationVariable(t *testing.T) {
	var validate *validator.Validate = validator.New()
	user := "tag"

	err := validate.Var(user, "required")
	if err != nil {
		t.Error(err)
	}
}
