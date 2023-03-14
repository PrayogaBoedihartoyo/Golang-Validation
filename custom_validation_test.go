package Golang_Validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
	"testing"
)

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}
	return true
}

func TestMustValidUsername(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("username", MustValidUsername)
	if err != nil {
		fmt.Println(err.Error())
	}

	type User struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	user := User{
		Username: "AABDWADC",
		Password: "123456",
	}
	err = validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}
