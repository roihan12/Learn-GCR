package recipes

import (
	"echo-recipe/businesses/recipes"

	"gorm.io/gorm"
)

type recipeRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) recipes.Repository {
	return &recipeRepository{
		conn: conn,
	}
}

func (re *recipeRepository) GetAll(name string) []recipes.Domain {
	var rec []Recipe

	re.conn.Preload("User").Preload("Category").Find(&rec)

	if name != "" {
		re.conn.Preload("User").Preload("Category").Where("name LIKE ?", "%"+name+"%").Find(&rec)
	}

	recipeDomain := []recipes.Domain{}

	for _, recipe := range rec {
		recipeDomain = append(recipeDomain, recipe.ToDomain())
	}

	return recipeDomain
}

func (re *recipeRepository) GetByID(id string) recipes.Domain {
	var recipe Recipe

	re.conn.Preload("User").Preload("Category").First(&recipe, "id = ?", id)

	return recipe.ToDomain()
}

func (re *recipeRepository) GetByCategoryID(categoryID string) []recipes.Domain {
	var rec []Recipe

	re.conn.Preload("User").Preload("Category").Find(&rec, "category_id = ?", categoryID)

	recipeDomain := []recipes.Domain{}

	for _, recipe := range rec {
		recipeDomain = append(recipeDomain, recipe.ToDomain())
	}

	return recipeDomain
}

func (re *recipeRepository) Create(recipeDomain *recipes.Domain) recipes.Domain {
	rec := FromDomain(recipeDomain)

	result := re.conn.Create(&rec)
	re.conn.Preload("User").Preload("Category").Find(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (re *recipeRepository) Update(id string, recipeDomain *recipes.Domain) recipes.Domain {
	var recipe recipes.Domain = re.GetByID(id)

	updatedRecipe := FromDomain(&recipe)

	updatedRecipe.Name = recipeDomain.Name
	updatedRecipe.Description = recipeDomain.Description
	updatedRecipe.Ingredients = recipeDomain.Ingredients
	updatedRecipe.Instructions = recipe.Instructions
	updatedRecipe.Difficult = recipeDomain.Difficult
	updatedRecipe.Time = recipeDomain.Time
	updatedRecipe.Serving = recipeDomain.Serving
	updatedRecipe.UserID = recipeDomain.UserID
	updatedRecipe.CategoryID = recipeDomain.CategoryID

	re.conn.Save(&updatedRecipe)
	re.conn.Preload("User").Preload("Category").Find(&updatedRecipe)
	return updatedRecipe.ToDomain()
}

func (re *recipeRepository) Delete(id string) bool {
	var recipe recipes.Domain = re.GetByID(id)

	deletedRecipe := FromDomain(&recipe)

	result := re.conn.Delete(&deletedRecipe)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
