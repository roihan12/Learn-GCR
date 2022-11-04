package controller

import (
	"echo-recipe/dto"
	"echo-recipe/entity"
	"echo-recipe/helper"
	"echo-recipe/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type RecipeController interface {
	All(ctx echo.Context) error
	FindByID(ctx echo.Context) error
	Insert(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type recipeController struct {
	recipeService service.RecipeService
	jwtService    service.JWTService
}

func NewRecipeController(recipe service.RecipeService, jwtServ service.JWTService) RecipeController {
	return &recipeController{
		recipeService: recipe,
		jwtService:    jwtServ,
	}
}

func (c *recipeController) All(ctx echo.Context) error {

	keyword := ctx.QueryParam("keyword")

	var recipes []entity.Recipe = c.recipeService.All(keyword)
	res := helper.BuildResponse(true, "all recipe", recipes)

	return ctx.JSON(http.StatusOK, res)
}

func (c *recipeController) FindByID(ctx echo.Context) error {
	var id string = ctx.Param("id")

	if id == "" {
		res := helper.BuildErrorResponse("No param id was found", "recipe not found", helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, res)
	}

	var recipe entity.Recipe = c.recipeService.FindByID(id)

	if (recipe == entity.Recipe{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		return ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "recipe found", recipe)
		return ctx.JSON(http.StatusOK, res)
	}
}

func (c *recipeController) Insert(ctx echo.Context) error {
	var recipeCreateDTO dto.RecipeDTO
	err := ctx.Bind(&recipeCreateDTO)

	if err := ctx.Validate(&recipeCreateDTO); err != nil {
		return err
	}

	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	} else {

		autheader := ctx.Request().Header.Get("Authorization")
		jwtString := strings.Split(autheader, "Bearer ")[1]
		userID := c.getUserIDByToken(jwtString)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			recipeCreateDTO.UserID = convertedUserID
		}
		result := c.recipeService.Insert(recipeCreateDTO)
		response := helper.BuildResponse(true, "recipe create", result)
		return ctx.JSON(http.StatusCreated, response)
	}

}

func (c *recipeController) Update(ctx echo.Context) error {

	var id string = ctx.Param("id")

	var recipeUpdateDTO dto.RecipeDTO
	err := ctx.Bind(&recipeUpdateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, res)

	}

	if err := ctx.Validate(&recipeUpdateDTO); err != nil {
		return err
	}

	autheader := ctx.Request().Header.Get("Authorization")
	jwtString := strings.Split(autheader, "Bearer ")[1]
	token, errToken := c.jwtService.ValidateToken(jwtString)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	convertedUserID, err := strconv.ParseUint(userID, 10, 64)

	recipeUpdateDTO.UserID = convertedUserID

	result := c.recipeService.Update(id, recipeUpdateDTO)

	response := helper.BuildResponse(true, "recipe update", result)
	return ctx.JSON(http.StatusOK, response)

}

func (c *recipeController) Delete(ctx echo.Context) error {
	recipeId := ctx.Param("id")

	isSucces := c.recipeService.Delete(recipeId)

	if !isSucces {
		res := helper.BuildErrorResponse("Failed to process request", "failed delete", helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, res)
	}

	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	return ctx.JSON(http.StatusOK, res)
}

func (c *recipeController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}

	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id

}
