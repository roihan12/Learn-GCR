package recipes

import (
	"echo-recipe/drivers/mysql/categories"
	"echo-recipe/drivers/mysql/users"
	"time"

	recipeUsecase "echo-recipe/businesses/recipes"

	"gorm.io/gorm"
)

type Recipe struct {
	ID           uint                `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time           `json:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at"`
	DeletedAt    gorm.DeletedAt      `json:"deleted_at"`
	Name         string              `gorm:"type:varchar(255)" json:"name"`
	Description  string              `json:"description"`
	Ingredients  string              `json:"ingredients"`
	Instructions string              `json:"instructions"`
	Difficult    string              `json:"difficult"`
	Time         string              `json:"time"`
	Serving      string              `json:"serving"`
	UserID       uint                `json:"user_id"`
	CategoryID   uint                `json:"category_id"`
	Category     categories.Category `json:"category" gorm:"foreignKey:CategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	User         users.User          `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"publisher"`
}

func FromDomain(domain *recipeUsecase.Domain) *Recipe {
	return &Recipe{
		ID:           domain.ID,
		Name:         domain.Name,
		Description:  domain.Description,
		Ingredients:  domain.Ingredients,
		Instructions: domain.Instructions,
		Difficult:    domain.Difficult,
		Time:         domain.Time,
		Serving:      domain.Serving,
		UserID:       domain.UserID,
		CategoryID:   domain.CategoryID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
	}
}

func (rec *Recipe) ToDomain() recipeUsecase.Domain {
	return recipeUsecase.Domain{
		ID:           rec.ID,
		Name:         rec.Name,
		Description:  rec.Description,
		Ingredients:  rec.Ingredients,
		Instructions: rec.Instructions,
		Difficult:    rec.Difficult,
		Time:         rec.Time,
		Serving:      rec.Serving,
		UserID:       rec.UserID,
		CategoryID:   rec.Category.ID,
		Category:     rec.Category.Name,
		User:         rec.User.Name,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
	}
}
