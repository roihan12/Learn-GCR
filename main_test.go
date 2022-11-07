package main

import (
	_driverFactory "echo-recipe/drivers"
	"echo-recipe/helper"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	_recipeUseCase "echo-recipe/businesses/recipes"
	_recipeController "echo-recipe/controllers/recipes"
	"echo-recipe/controllers/users/request"

	_categoryUseCase "echo-recipe/businesses/categories"
	_categoryController "echo-recipe/controllers/categories"

	_userUseCase "echo-recipe/businesses/users"
	_userController "echo-recipe/controllers/users"

	_dbDriver "echo-recipe/drivers/mysql"
	"echo-recipe/drivers/mysql/categories"
	"echo-recipe/drivers/mysql/recipes"
	"echo-recipe/drivers/mysql/users"

	_middleware "echo-recipe/app/middlewares"
	_routes "echo-recipe/app/routes"

	"github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/steinfletcher/apitest"
)

func main_test() *echo.Echo {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: helper.GetConfig("DB_USERNAME"),
		DB_PASSWORD: helper.GetConfig("DB_PASSWORD"),
		DB_HOST:     helper.GetConfig("DB_HOST"),
		DB_PORT:     helper.GetConfig("DB_PORT"),
		DB_NAME:     helper.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.DBMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       helper.GetConfig("JWT_SECRET_KEY"),
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

	return e
}

func cleanup(res *http.Response, req *http.Request, apiTest *apitest.APITest) {
	if http.StatusOK == res.StatusCode || http.StatusCreated == res.StatusCode {
		configDB := _dbDriver.ConfigDB{
			DB_USERNAME: helper.GetConfig("DB_USERNAME"),
			DB_PASSWORD: helper.GetConfig("DB_PASSWORD"),
			DB_HOST:     helper.GetConfig("DB_HOST"),
			DB_PORT:     helper.GetConfig("DB_PORT"),
			DB_NAME:     helper.GetConfig("DB_TEST_NAME"),
		}

		db := configDB.InitDB()

		_dbDriver.CleanSeeders(db)
	}
}

func getJWTToken(t *testing.T) string {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: helper.GetConfig("DB_USERNAME"),
		DB_PASSWORD: helper.GetConfig("DB_PASSWORD"),
		DB_HOST:     helper.GetConfig("DB_HOST"),
		DB_PORT:     helper.GetConfig("DB_PORT"),
		DB_NAME:     helper.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	user := _dbDriver.SeedUser(db)

	var userRequest *request.User = &request.User{
		Email:    user.Email,
		Password: user.Password,
	}

	var resp *http.Response = apitest.New().
		Handler(main_test()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End().Response

	var response map[string]string = map[string]string{}

	json.NewDecoder(resp.Body).Decode(&response)

	var token string = response["token"]

	var JWT_TOKEN = "Bearer " + token

	return JWT_TOKEN
}

func getUser() users.User {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: helper.GetConfig("DB_USERNAME"),
		DB_PASSWORD: helper.GetConfig("DB_PASSWORD"),
		DB_HOST:     helper.GetConfig("DB_HOST"),
		DB_PORT:     helper.GetConfig("DB_PORT"),
		DB_NAME:     helper.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	user := _dbDriver.SeedUser(db)

	return user
}

func getRecipe() recipes.Recipe {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: helper.GetConfig("DB_USERNAME"),
		DB_PASSWORD: helper.GetConfig("DB_PASSWORD"),
		DB_HOST:     helper.GetConfig("DB_HOST"),
		DB_PORT:     helper.GetConfig("DB_PORT"),
		DB_NAME:     helper.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	recipe := _dbDriver.SeedRecipe(db)

	return recipe
}

func getCategory() categories.Category {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: helper.GetConfig("DB_USERNAME"),
		DB_PASSWORD: helper.GetConfig("DB_PASSWORD"),
		DB_HOST:     helper.GetConfig("DB_HOST"),
		DB_PORT:     helper.GetConfig("DB_PORT"),
		DB_NAME:     helper.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	category := _dbDriver.SeedCategory(db)

	return category
}

func TestRegister_Success(t *testing.T) {
	var userRequest *request.User = &request.User{
		Name:     "testname",
		Email:    "test12@mail.com",
		Password: "123123",
	}

	apitest.New().
		Observe(cleanup).
		Handler(main_test()).
		Post("/api/v1/users/register").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestRegister_ValidationFailed(t *testing.T) {
	var userRequest *request.User = &request.User{
		Name:     "",
		Email:    "",
		Password: "",
	}

	apitest.New().
		Handler(main_test()).
		Post("/api/v1/users/register").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Success(t *testing.T) {
	user := getUser()

	var userRequest *request.User = &request.User{
		Email:    user.Email,
		Password: user.Password,
	}

	apitest.New().
		Handler(main_test()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestLogin_ValidationFailed(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "",
		Password: "",
	}

	apitest.New().
		Handler(main_test()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Failed(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "notfound@mail.com",
		Password: "123123",
	}

	apitest.New().
		Handler(main_test()).
		Post("/api/v1/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusUnauthorized).
		End()
}

func TestGetRecipe_Success(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(main_test()).
		Get("/api/v1/recipes").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetRecipes_Success(t *testing.T) {
	var recipe recipes.Recipe = getRecipe()

	recipeID := strconv.Itoa(int(recipe.ID))

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(main_test()).
		Get("/api/v1/recipes/"+recipeID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetRecipes_NotFound(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Handler(main_test()).
		Get("/api/v1/recipes/0").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestCreateRecipe_Success(t *testing.T) {
	category := getCategory()
	user := getUser()

	var recipeRequest *recipes.Recipe = &recipes.Recipe{
		Name:         "testing",
		Description:  "ini desc",
		Ingredients:  "test ingredients",
		Instructions: "test insructions",
		Difficult:    "mudah",
		Time:         "4 jam",
		Serving:      "5 porsi",
		UserID:       user.ID,
		CategoryID:   category.ID,
	}

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(main_test()).
		Post("/api/v1/recipes").
		Header("Authorization", token).
		JSON(recipeRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestCreateRecipe_ValidationFailed(t *testing.T) {
	var recipeRequest *recipes.Recipe = &recipes.Recipe{}

	var token string = getJWTToken(t)

	apitest.New().
		Handler(main_test()).
		Post("/api/v1/recipes").
		Header("Authorization", token).
		JSON(recipeRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestUpdateRecipe_Success(t *testing.T) {

	var recipe recipes.Recipe = getRecipe()
	category := getCategory()
	user := getUser()

	recipeID := strconv.Itoa(int(recipe.ID))
	var recipeRequest *recipes.Recipe = &recipes.Recipe{
		Name:         "testing updated",
		Description:  "ini desc",
		Ingredients:  "test ingredients",
		Instructions: "test insructions",
		Difficult:    "mudah",
		Time:         "4 jam",
		Serving:      "5 porsi",
		UserID:       user.ID,
		CategoryID:   category.ID,
	}

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(main_test()).
		Put("/api/v1/recipes/"+recipeID).
		Header("Authorization", token).
		JSON(recipeRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestUpdateNote_ValidationFailed(t *testing.T) {
	var recipe recipes.Recipe = getRecipe()

	var recipeRequest *recipes.Recipe = &recipes.Recipe{}

	recipeID := strconv.Itoa(int(recipe.ID))

	var token string = getJWTToken(t)

	apitest.New().
		Handler(main_test()).
		Put("/api/v1/recipes/"+recipeID).
		Header("Authorization", token).
		JSON(recipeRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestDeleteRecipe_Success(t *testing.T) {
	var recipe recipes.Recipe = getRecipe()

	var token string = getJWTToken(t)

	recipeID := strconv.Itoa(int(recipe.ID))

	apitest.New().
		Observe(cleanup).
		Handler(main_test()).
		Delete("/api/v1/recipes/"+recipeID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestDeleteRecipe_Failed(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Handler(main_test()).
		Observe(cleanup).
		Delete("/api/v1/recipe/-1").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestLogout_Success(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Handler(main_test()).
		Observe(cleanup).
		Post("/api/v1/users/logout").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestLogout_Failed(t *testing.T) {
	apitest.New().
		Handler(main_test()).
		Observe(cleanup).
		Post("/api/v1/users/logout").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}
