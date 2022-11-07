package drivers

import (
	categoryDomain "echo-recipe/businesses/categories"
	categoryDB "echo-recipe/drivers/mysql/categories"

	recipeDomain "echo-recipe/businesses/recipes"
	recipeDB "echo-recipe/drivers/mysql/recipes"

	userDomain "echo-recipe/businesses/users"
	userDB "echo-recipe/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLRepository(conn)
}

func NewNoteRepository(conn *gorm.DB) recipeDomain.Repository {
	return recipeDB.NewMySQLRepository(conn)
}

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}
