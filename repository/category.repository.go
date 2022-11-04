package repository

import (
	"echo-recipe/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Insert(category entity.Category) entity.Category
	Update(id string, category entity.Category) entity.Category
	Delete(id string) bool
	All() []entity.Category
	FindByID(categoryID string) entity.Category
}

type categoryConnection struct {
	connection *gorm.DB
}

func NewCategoryRepository(dbConn *gorm.DB) CategoryRepository {
	return &categoryConnection{
		connection: dbConn,
	}
}

func (db *categoryConnection) Insert(category entity.Category) entity.Category {
	db.connection.Save(&category)
	db.connection.Find(&category)
	return category
}

func (db *categoryConnection) Update(id string, input entity.Category) entity.Category {

	var category entity.Category = db.FindByID(id)

	category.Name = input.Name

	db.connection.Save(&category)
	db.connection.Find(&category)

	return category
}

func (db *categoryConnection) Delete(id string) bool {

	var category entity.Category
	res := db.connection.Where("id = ?", id).Take(&category)
	db.connection.Delete(&category)

	if res.RowsAffected == 0 {
		return false
	}

	return true

}

func (db *categoryConnection) FindByID(id string) entity.Category {
	var category entity.Category
	db.connection.Preload("Recipe").Find(&category, id)
	return category
}

func (db *categoryConnection) All() []entity.Category {
	var category []entity.Category
	db.connection.Find(&category)
	return category
}
