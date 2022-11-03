package routes

import (
	"echo-recipe/controller"
	"echo-recipe/database"

	"echo-recipe/middlewares"
	"echo-recipe/repository"
	"echo-recipe/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	db                 *gorm.DB                      = database.SetupDatabaseConnection()
	userRepository     repository.UserRepository     = repository.NewUserRepository(db)
	recipeRepository   repository.RecipeRepository   = repository.NewRecipeRepository(db)
	categoryRepository repository.CategoryRepository = repository.NewCategoryRepository(db)
	jwtService         service.JWTService            = service.NewJWTService()
	authService        service.AuthService           = service.NewAuthService(userRepository)
	userService        service.UserService           = service.NewUserService(userRepository)
	recipeService      service.RecipeService         = service.NewRecipeService(recipeRepository)
	categoryService    service.CategoryService       = service.NewCategoryService(categoryRepository)
	authController     controller.AuthController     = controller.NewAuthController(authService, jwtService)
	userController     controller.UserController     = controller.NewUserController(userService, jwtService)
	recipeController   controller.RecipeController   = controller.NewRecipeController(recipeService, jwtService)
	categoryController controller.CategoryController = controller.NewCategoryController(categoryService)
)

func SetupRoute(server *echo.Echo) {
	// defer database.CloseDatabaseConnection(db)

	authRoute := server.Group("api/v1/auth")
	authRoute.POST("/login", authController.Login)
	authRoute.POST("/register", authController.Register)

	userRoute := server.Group("api/v1/users", middlewares.AuthorizeJWT)
	userRoute.GET("/profile", userController.Profile)
	userRoute.PUT("/profile", userController.Update)

	recipeRoute := server.Group("api/v1/recipe", middlewares.AuthorizeJWT)
	recipeRoute.GET("", recipeController.All)
	recipeRoute.POST("", recipeController.Insert)
	recipeRoute.GET("/:id", recipeController.FindByID)
	recipeRoute.PUT("/:id", recipeController.Update)
	recipeRoute.DELETE("/:id", recipeController.Delete)

	// routes for categories
	categoriesRoute := server.Group("api/v1/categories", middlewares.AuthorizeJWT)
	categoriesRoute.GET("", categoryController.All)
	categoriesRoute.POST("", categoryController.Insert)
	categoriesRoute.GET("/:id", categoryController.FindByID)
	categoriesRoute.PUT("/:id", categoryController.Update)
	categoriesRoute.DELETE("/:id", categoryController.Delete)
}
