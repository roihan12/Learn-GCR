package database

import (
	"echo-recipe/entity"
	"echo-recipe/helper"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB_USERNAME string = helper.GetConfig("DB_USERNAME")
	DB_PASSWORD string = helper.GetConfig("DB_PASSWORD")
	DB_NAME     string = helper.GetConfig("DB_NAME")
	DB_HOST     string = helper.GetConfig("DB_HOST")
	DB_PORT     string = helper.GetConfig("DB_PORT")
)

// SetupDatabaseConnection berfungsi  untuk koneksi ke database
func SetupDatabaseConnection() *gorm.DB {

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to create a connection to database")
	}

	//Membuat model database
	db.AutoMigrate(&entity.User{}, &entity.Recipe{}, &entity.Category{})
	return db
}

// CloseDatabaseConnection berfungsi untuk menutup koneksi database kita
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("failed to close connection database")
	}

	dbSQL.Close()

}
