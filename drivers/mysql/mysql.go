package mysql_driver

import (
	"echo-recipe/drivers/mysql/categories"
	"echo-recipe/drivers/mysql/recipes"
	"echo-recipe/drivers/mysql/users"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"

	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func (config *ConfigDB) InitDB() *gorm.DB {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")

	return db
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&recipes.Recipe{}, &categories.Category{}, &users.User{})
}

func CloseDB(db *gorm.DB) error {
	database, err := db.DB()

	if err != nil {
		log.Printf("error when getting the database instance: %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing the database connection: %v", err)
		return err
	}

	log.Println("database connection is closed")

	return nil
}

func SeedUser(db *gorm.DB) users.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)

	var user users.User = users.User{
		Name:     "Test1",
		Email:    "testin1g@mail.com",
		Password: string(password),
	}

	db.Create(&user)

	var createdUser users.User

	db.Last(&createdUser)

	createdUser.Password = "123123"

	return createdUser
}

func SeedRecipe(db *gorm.DB) recipes.Recipe {
	category := SeedCategory(db)

	user := SeedUser(db)

	var recipe recipes.Recipe = recipes.Recipe{
		Name:         "test",
		Description:  "ini desc",
		Ingredients:  "test ingredients",
		Instructions: "test insructions",
		Difficult:    "mudah",
		Time:         "4 jam",
		Serving:      "5 porsi",
		UserID:       user.ID,
		CategoryID:   category.ID,
	}

	if err := db.Create(&recipe).Error; err != nil {
		panic(err)
	}

	var createdRecipe recipes.Recipe

	db.Last(&createdRecipe)

	return createdRecipe
}

func SeedCategory(db *gorm.DB) categories.Category {

	var category categories.Category = categories.Category{
		Name: "test",
	}

	if err := db.Create(&category).Error; err != nil {
		panic(err)
	}

	var createdCategory categories.Category

	db.Last(&createdCategory)

	return createdCategory
}

func CleanSeeders(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	categoryResult := db.Exec("DELETE FROM categories")
	itemResult := db.Exec("DELETE FROM recipes")
	userResult := db.Exec("DELETE FROM users")

	var isFailed bool = itemResult.Error != nil || userResult.Error != nil || categoryResult.Error != nil
	if isFailed {
		panic(errors.New("error when cleaning up seeders"))
	}

	log.Println("Seeders are cleaned up successfully")
}
