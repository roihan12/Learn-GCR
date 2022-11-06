package recipes

import (
	"echo-recipe/app/middlewares"
	"echo-recipe/businesses/recipes"
	controller "echo-recipe/controllers"
	"echo-recipe/controllers/recipes/request"
	"echo-recipe/controllers/recipes/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RecipeController struct {
	recipeUseCase recipes.Usecase
}

func NewNoteController(recipeUC recipes.Usecase) *RecipeController {
	return &RecipeController{
		recipeUseCase: recipeUC,
	}
}

func (ctrl *RecipeController) GetAll(c echo.Context) error {
	recipesData := ctrl.recipeUseCase.GetAll()

	recipes := []response.RecipeAll{}

	for _, recipe := range recipesData {
		recipes = append(recipes, response.FromDomainGetAll(recipe))
	}

	return controller.NewResponse(c, http.StatusOK, "success", "all recipes", recipes)
}

func (ctrl *RecipeController) GetByID(c echo.Context) error {
	var id string = c.Param("id")

	recipe := ctrl.recipeUseCase.GetByID(id)

	if recipe.ID == 0 {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "recipe not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "recipe found", response.FromDomain(recipe))
}

func (ctrl *RecipeController) Create(c echo.Context) error {

	input := request.Recipe{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	if err := c.Validate(&input); err != nil {
		return err
	}

	user := middlewares.GetUserID(c)
	input.UserID = user.ID

	recipe := ctrl.recipeUseCase.Create(input.ToDomain())

	return controller.NewResponse(c, http.StatusCreated, "success", "recipe created", response.FromDomain(recipe))
}

func (ctrl *RecipeController) Update(c echo.Context) error {
	input := request.Recipe{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	var recipeId string = c.Param("id")

	if err := c.Validate(&input); err != nil {
		return err
	}

	user := middlewares.GetUserID(c)
	input.UserID = user.ID

	recipe := ctrl.recipeUseCase.Update(recipeId, input.ToDomain())

	if recipe.ID == 0 {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "recipe not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "recipe updated", response.FromDomain(recipe))
}

func (ctrl *RecipeController) Delete(c echo.Context) error {
	var recipeId string = c.Param("id")

	isSuccess := ctrl.recipeUseCase.Delete(recipeId)

	if !isSuccess {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "recipe not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "recipe deleted", "")
}
