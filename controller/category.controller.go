package controller

import (
	"echo-recipe/dto"
	"echo-recipe/entity"
	"echo-recipe/helper"
	"echo-recipe/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController interface {
	All(ctx echo.Context) error
	FindByID(ctx echo.Context) error
	Insert(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type categoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(category service.CategoryService) CategoryController {
	return &categoryController{
		categoryService: category,
	}
}

func (c *categoryController) All(ctx echo.Context) error {
	var categories []entity.Category = c.categoryService.All()

	res := helper.BuildResponse(true, "all category", categories)

	return ctx.JSON(http.StatusOK, res)
}

func (c *categoryController) FindByID(ctx echo.Context) error {
	var id string = ctx.Param("id")

	if id == "" {
		res := helper.BuildErrorResponse("No param id was found", "not found", helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, res)
	}

	var categories entity.Category = c.categoryService.FindByID(id)

	if (categories == entity.Category{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		return ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "category found", categories)
		return ctx.JSON(http.StatusOK, res)
	}
}

func (c *categoryController) Insert(ctx echo.Context) error {
	var categoriesCreate dto.CategoryDTO
	err := ctx.Bind(&categoriesCreate)

	if err := ctx.Validate(&categoriesCreate); err != nil {
		return err
	}

	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	result := c.categoryService.Insert(categoriesCreate)
	response := helper.BuildResponse(true, "category create", result)
	return ctx.JSON(http.StatusCreated, response)

}

func (c *categoryController) Update(ctx echo.Context) error {

	var id string = ctx.Param("id")

	var categoriesUpdateDTO dto.CategoryDTO
	err := ctx.Bind(&categoriesUpdateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, res)

	}

	if err := ctx.Validate(&categoriesUpdateDTO); err != nil {
		return err
	}

	result := c.categoryService.Update(id, categoriesUpdateDTO)

	response := helper.BuildResponse(true, "catogory update", result)
	return ctx.JSON(http.StatusOK, response)

}

func (c *categoryController) Delete(ctx echo.Context) error {
	categoryId := ctx.Param("id")

	isSucces := c.categoryService.Delete(categoryId)

	if !isSucces {
		res := helper.BuildErrorResponse("Failed to process request", "failed delete", helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, res)
	}

	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	return ctx.JSON(http.StatusOK, res)
}
