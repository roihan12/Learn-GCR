package recipes_test

import (
	"echo-recipe/businesses/categories"
	"echo-recipe/businesses/recipes"
	_recipeMock "echo-recipe/businesses/recipes/mocks"
	"echo-recipe/businesses/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	recipeRepository _recipeMock.Repository
	recipeService    recipes.Usecase

	recipeDomain recipes.Domain
)

func TestMain(m *testing.M) {
	recipeService = recipes.NewNoteUsecase(&recipeRepository)

	categoryDomain := categories.Domain{
		Name: "test category",
	}

	usersDomain := users.Domain{
		Email:    "test@gmail.com",
		Password: "tri123",
	}
	recipeDomain = recipes.Domain{
		Name:         "test name",
		Description:  "test desc",
		Ingredients:  "test ingre",
		Instructions: "test instruc",
		Difficult:    "test dif",
		Time:         "4 jam",
		Serving:      "5 porsi",
		UserID:       usersDomain.ID,
		CategoryID:   categoryDomain.ID,
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		recipeRepository.On("GetAll").Return([]recipes.Domain{recipeDomain}).Once()

		result := recipeService.GetAll()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		recipeRepository.On("GetAll").Return([]recipes.Domain{}).Once()

		result := recipeService.GetAll()

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Get By ID | Valid", func(t *testing.T) {
		recipeRepository.On("GetByID", "1").Return(recipeDomain).Once()

		result := recipeService.GetByID("1")

		assert.NotNil(t, result)
	})

	t.Run("Get By ID | InValid", func(t *testing.T) {
		recipeRepository.On("GetByID", "-1").Return(recipes.Domain{}).Once()

		result := recipeService.GetByID("-1")

		assert.NotNil(t, result)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		recipeRepository.On("Create", &recipeDomain).Return(recipeDomain).Once()

		result := recipeService.Create(&recipeDomain)

		assert.NotNil(t, result)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		recipeRepository.On("Create", &recipes.Domain{}).Return(recipes.Domain{}).Once()

		result := recipeService.Create(&recipes.Domain{})

		assert.NotNil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		recipeRepository.On("Update", "1", &recipeDomain).Return(recipeDomain).Once()

		result := recipeService.Update("1", &recipeDomain)

		assert.NotNil(t, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		recipeRepository.On("Update", "1", &recipes.Domain{}).Return(recipes.Domain{}).Once()

		result := recipeService.Update("1", &recipes.Domain{})

		assert.NotNil(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		recipeRepository.On("Delete", "1").Return(true).Once()

		result := recipeService.Delete("1")

		assert.True(t, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		recipeRepository.On("Delete", "-1").Return(false).Once()

		result := recipeService.Delete("-1")

		assert.False(t, result)
	})
}
