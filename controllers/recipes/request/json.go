package request

import (
	"echo-recipe/businesses/recipes"

	"github.com/go-playground/validator/v10"
)

type Recipe struct {
	Name         string `json:"name" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Ingredients  string `json:"ingredients" validate:"required"`
	Instructions string `json:"instructions" validate:"required"`
	Difficult    string `json:"difficult" validate:"required"`
	Time         string `json:"time" validate:"required"`
	Serving      string `json:"serving" validate:"required"`
	CategoryID   uint   `json:"category_id" validate:"required"`
	UserID       uint   `json:"user_id,omitempty"`
}

func (req *Recipe) ToDomain() *recipes.Domain {
	return &recipes.Domain{
		Name:         req.Name,
		Description:  req.Description,
		Ingredients:  req.Ingredients,
		Instructions: req.Instructions,
		Difficult:    req.Difficult,
		Time:         req.Time,
		Serving:      req.Serving,
		CategoryID:   req.CategoryID,
		UserID:       req.UserID,
	}
}

func (input *Recipe) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
