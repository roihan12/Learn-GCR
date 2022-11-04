package repository

import (
	"echo-recipe/entity"
	"fmt"

	"gorm.io/gorm"
)

type RecipeRepository interface {
	InsertRecipe(recipe entity.Recipe) entity.Recipe
	UpdateRecipe(id string, recipe entity.Recipe) entity.Recipe
	DeleteRecipe(recipeID string) bool
	AllRecipe(keyword string) []entity.Recipe
	FindRecipeByID(recipeID string) entity.Recipe
}

type recipeConnection struct {
	connection *gorm.DB
}

func NewRecipeRepository(dbConn *gorm.DB) RecipeRepository {
	return &recipeConnection{
		connection: dbConn,
	}
}

func (db *recipeConnection) InsertRecipe(recipe entity.Recipe) entity.Recipe {
	db.connection.Save(&recipe)
	db.connection.Preload("User").Preload("Category").Find(&recipe)
	return recipe
}

func (db *recipeConnection) UpdateRecipe(id string, input entity.Recipe) entity.Recipe {

	var recipe entity.Recipe = db.FindRecipeByID(id)

	recipe.Name = input.Name
	recipe.Description = input.Description
	recipe.Ingredients = input.Ingredients
	recipe.Instructions = input.Instructions
	recipe.Time = input.Time
	recipe.Serving = input.Serving
	recipe.Category = input.Category
	recipe.UserID = input.UserID

	db.connection.Save(&recipe)
	db.connection.Preload("Category").Preload("User").Find(&recipe)

	return recipe
}

func (db *recipeConnection) DeleteRecipe(recipeID string) bool {

	var recipe entity.Recipe
	res := db.connection.Preload("User").Where("id = ?", recipeID).Take(&recipe)
	db.connection.Delete(&recipe)

	if res.RowsAffected == 0 {
		return false
	}

	return true

}

func (db *recipeConnection) FindRecipeByID(recipeID string) entity.Recipe {
	var recipe entity.Recipe
	db.connection.Preload("Category").Preload("User").Find(&recipe, recipeID)
	return recipe
}

func (db *recipeConnection) AllRecipe(keyword string) []entity.Recipe {
	var recipe []entity.Recipe

	sql := "SELECT * FROM recipes"
	s := keyword
	if s != "" {
		sql = fmt.Sprintf("%s WHERE name LIKE '%%%s%%' OR description LIKE '%%%s%%'", sql, s, s)

	}

	db.connection.Debug().Preload("Category").Preload("User").Raw(sql).Find(&recipe)
	return recipe
}
