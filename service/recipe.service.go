package service

import (
	"echo-recipe/dto"
	"echo-recipe/entity"
	"echo-recipe/repository"
	"log"

	"github.com/mashingan/smapping"
)

// BookService is a ....
type RecipeService interface {
	Insert(b dto.RecipeDTO) entity.Recipe
	Update(id string, b dto.RecipeDTO) entity.Recipe
	Delete(id string) bool
	All(keyword string) []entity.Recipe
	FindByID(recipekID string) entity.Recipe
	FindByCategoryID(categoryId string) []entity.Recipe
}

type recipeService struct {
	recipeRepository repository.RecipeRepository
}

// NewBookService .....
func NewRecipeService(recipeRepo repository.RecipeRepository) RecipeService {
	return &recipeService{
		recipeRepository: recipeRepo,
	}
}

func (service *recipeService) Insert(b dto.RecipeDTO) entity.Recipe {

	recipe := entity.Recipe{}
	err := smapping.FillStruct(&recipe, smapping.MapFields(&b))

	if err != nil {
		log.Fatalf("failed map %v", err)
	}

	res := service.recipeRepository.InsertRecipe(recipe)
	return res

}

func (service *recipeService) Update(id string, b dto.RecipeDTO) entity.Recipe {

	recipe := entity.Recipe{}
	err := smapping.FillStruct(&recipe, smapping.MapFields(&b))

	if err != nil {
		log.Fatalf("failed map %v", err)
	}

	res := service.recipeRepository.UpdateRecipe(id, recipe)
	return res

}

func (service *recipeService) Delete(id string) bool {
	return service.recipeRepository.DeleteRecipe(id)

}

func (service *recipeService) All(keyword string) []entity.Recipe {
	return service.recipeRepository.AllRecipe(keyword)

}

func (service *recipeService) FindByID(recipeID string) entity.Recipe {
	return service.recipeRepository.FindRecipeByID(recipeID)

}

func (service *recipeService) FindByCategoryID(categoryId string) []entity.Recipe {
	return service.recipeRepository.FindRecipeByCategoryID(categoryId)
}
