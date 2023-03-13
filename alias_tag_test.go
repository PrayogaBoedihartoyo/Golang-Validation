package Golang_Validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestAliasTag(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("alias", "required")
	type User struct {
		Name string `validate:"alias"`
	}
	user := User{
		Name: "",
	}
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}
