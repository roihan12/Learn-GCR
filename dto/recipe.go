package dto

import "github.com/go-playground/validator/v10"

type RecipeDTO struct {
	Name         string `json:"name" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Ingredients  string `json:"ingredients" validate:"required"`
	Instructions string `json:"instructions" validate:"required"`
	Time         string `json:"time" validate:"required"`
	Serving      string `json:"serving" validate:"required"`
	CategoryID   uint64 `json:"category_id" validate:"required"`
	UserID       uint64 `json:"user_id,omitempty"`
}

func (input *RecipeDTO) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
