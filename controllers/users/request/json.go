package request

import (
	"echo-recipe/businesses/users"

	"github.com/go-playground/validator/v10"
)

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (req *User) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *UserLogin) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *User) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

func (req *UserLogin) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
