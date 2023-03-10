package Golang_Validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestValidationStruct(t *testing.T) {
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
		fmt.Println(err.Error())
	}

}

func TestNestedStruct(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}

	validate := validator.New()
	loginRequest := User{
		Name: "Prayoga",
		Address: Address{
			City:    "Jakarta",
			Country: "Indonesia",
		},
	}
	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
	}

	validate := validator.New()
	loginRequest := User{
		Name: "Prayoga",
		Addresses: []Address{
			{
				City:    "Jakarta",
				Country: "Indonesia",
			}, {
				City:    "Bandung",
				Country: "",
			},
		},
	}
	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestBasicCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
		Hobbies   []string  `validate:"dive,required,min=1"`
	}

	validate := validator.New()
	loginRequest := User{
		Name: "Prayoga",
		Addresses: []Address{
			{
				City:    "Jakarta",
				Country: "Indonesia",
			}, {
				City:    "Bandung",
				Country: "",
			},
		},
		Hobbies: []string{
			"Reading",
			"Writing",
			"",
		},
	}
	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}
