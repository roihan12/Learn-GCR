package main

import (
	_driverFactory "echo-recipe/drivers"
	"echo-recipe/helper"
	"fmt"
	"os"

	_recipeUseCase "echo-recipe/businesses/recipes"
	_recipeController "echo-recipe/controllers/recipes"

	_categoryUseCase "echo-recipe/businesses/categories"
	_categoryController "echo-recipe/controllers/categories"

	_userUseCase "echo-recipe/businesses/users"
	_userController "echo-recipe/controllers/users"

	_dbDriver "echo-recipe/drivers/mysql"

	_middleware "echo-recipe/app/middlewares"
	_routes "echo-recipe/app/routes"

	"github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const DEFAULT_PORT = "8080"

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: helper.GetConfig("DB_USERNAME"), //os.Getenv("DB_USERNAME"),  ,
		DB_PASSWORD: helper.GetConfig("DB_PASSWORD"), //os.Getenv("DB_PASSWORD"), ,
		DB_HOST:     helper.GetConfig("DB_HOST"),     //os.Getenv("DB_HOST"),
		DB_PORT:     helper.GetConfig("DB_PORT"),     //os.Getenv("DB_PORT"),
		DB_NAME:     helper.GetConfig("DB_NAME"),     //os.Getenv("DB_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.DBMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       os.Getenv("JWT_SECRETKEY"),
		ExpiresDuration: 1,
	}

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(_middleware.LoggerConfig()))
	e.Validator = &helper.CustomValidator{Validator: validator.New()}

	categoryRepo := _driverFactory.NewCategoryRepository(db)
	categoryUsecase := _categoryUseCase.NewCategoryUsecase(categoryRepo)
	categoryCtrl := _categoryController.NewCategoryController(categoryUsecase)

	recipeRepo := _driverFactory.NewNoteRepository(db)
	recipeUsecase := _recipeUseCase.NewNoteUsecase(recipeRepo)
	recipeCtrl := _recipeController.NewNoteController(recipeUsecase)

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUseCase.NewUserUsecase(userRepo, &configJWT)
	userCtrl := _userController.NewAuthController(userUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:      configJWT.Init(),
		CategoryController: *categoryCtrl,
		RecipeController:   *recipeCtrl,
		AuthController:     *userCtrl,
	}

	routesInit.RouteRegister(e)

	var port string = os.Getenv("PORT")

	if port == "" {
		port = DEFAULT_PORT
	}
	var appPort string = fmt.Sprintf(":%s", port)

	e.Logger.Fatal(e.Start(appPort))

}
