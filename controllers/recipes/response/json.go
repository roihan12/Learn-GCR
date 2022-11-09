package response

import (
	"echo-recipe/businesses/recipes"
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Ingredients  string         `json:"ingredients"`
	Instructions string         `json:"instructions"`
	Difficult    string         `json:"difficult"`
	Time         string         `json:"time"`
	Serving      string         `json:"serving"`
	UserID       uint           `json:"user_id"`
	CategoryID   uint           `json:"category_id"`
	Category     string         `json:"category"`
	User         string         `json:"publisher"`
}

type RecipeAll struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Ingredients  string         `json:"-"`
	Instructions string         `json:"-"`
	Difficult    string         `json:"difficult"`
	Time         string         `json:"time"`
	Serving      string         `json:"serving"`
	CategoryID   uint           `json:"-"`
	Category     string         `json:"category"`
	UserID       uint           `json:"-"`
	User         string         `json:"publisher"`
}

func FromDomainGetAll(domain recipes.Domain) RecipeAll {
	return RecipeAll{

		ID:          domain.ID,
		Name:        domain.Name,
		Description: domain.Description,
		Difficult:   domain.Difficult,
		Time:        domain.Time,
		Serving:     domain.Serving,
		UserID:      domain.UserID,
		User:        domain.User,
		CategoryID:  domain.CategoryID,
		Category:    domain.Category,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
	}
}

func FromDomain(domain recipes.Domain) Recipe {
	return Recipe{

		ID:           domain.ID,
		Name:         domain.Name,
		Description:  domain.Description,
		Ingredients:  domain.Ingredients,
		Instructions: domain.Instructions,
		Difficult:    domain.Difficult,
		Time:         domain.Time,
		Serving:      domain.Serving,
		UserID:       domain.UserID,
		User:         domain.User,
		CategoryID:   domain.CategoryID,
		Category:     domain.Category,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
	}
}
