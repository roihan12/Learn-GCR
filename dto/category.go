package dto

import "github.com/go-playground/validator/v10"

type CategoryDTO struct {
	Name string `json:"name" validate:"required"`
}

func (input *CategoryDTO) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
