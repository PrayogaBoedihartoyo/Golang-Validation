package Golang_Validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"testing"
)

var regexNumber = regexp.MustCompile(`^[0-9]+$`)

func MusValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}
	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}
	return len(value) == length
}

func TestValidPin(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("pin", MusValidPin)
	if err != nil {
		panic(err)
	}

	type Login struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}
	request := Login{
		Phone: "123456",
		Pin:   "123",
	}
	err = validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}
