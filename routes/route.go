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
	db             *gorm.DB                  = database.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	userService    service.UserService       = service.NewUserService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func SetupRoute(server *echo.Echo) {
	// defer database.CloseDatabaseConnection(db)

	authRoute := server.Group("api/v1/auth")
	authRoute.POST("/login", authController.Login)
	authRoute.POST("/register", authController.Register)

	userRoute := server.Group("api/v1/users", middlewares.AuthorizeJWT)
	userRoute.GET("/profile", userController.Profile)
	userRoute.PUT("/profile", userController.Update)
}
