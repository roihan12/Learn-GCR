package routes

import (
	"echo-recipe/controllers/categories"
	"echo-recipe/controllers/recipes"
	"echo-recipe/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	AuthController     users.AuthController
	RecipeController   recipes.RecipeController
	CategoryController categories.CategoryController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	users := e.Group("/api/v1/users")

	users.POST("/register", cl.AuthController.Register)
	users.POST("/login", cl.AuthController.Login)

	usersUpdate := e.Group("/api/v1/users", middleware.JWTWithConfig(cl.JWTMiddleware))
	usersUpdate.PUT("/update", cl.AuthController.Update)

	recipe := e.Group("/api/v1/recipes", middleware.JWTWithConfig(cl.JWTMiddleware))

	recipe.GET("", cl.RecipeController.GetAll)
	recipe.GET("/:id", cl.RecipeController.GetByID)
	recipe.POST("", cl.RecipeController.Create)
	recipe.PUT("/:id", cl.RecipeController.Update)
	recipe.DELETE("/:id", cl.RecipeController.Delete)

	category := e.Group("/api/v1/categories", middleware.JWTWithConfig(cl.JWTMiddleware))

	category.GET("", cl.CategoryController.GetAllCategories)
	category.POST("", cl.CategoryController.CreateCategory)
	category.PUT("/:id", cl.CategoryController.UpdateCategory)
	category.DELETE("/:id", cl.CategoryController.DeleteCategory)

	auth := e.Group("/api/v1/users", middleware.JWTWithConfig(cl.JWTMiddleware))

	auth.POST("/logout", cl.AuthController.Logout)

}
