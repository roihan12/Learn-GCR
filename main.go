// package main

// import (
// 	"echo-recipe/helper"
// 	"echo-recipe/middlewares"
// 	"echo-recipe/routes"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/labstack/echo/v4"
// )

// func main() {

// 	server := echo.New()

// 	server.Validator = &helper.CustomValidator{Validator: validator.New()}

// 	middlewares.LogMiddleware(server)

// 	routes.SetupRoute(server)

// 	server.Logger.Fatal(server.Start(":1323"))
// }

package main

import (
	"context"
	_driverFactory "echo-recipe/drivers"
	"echo-recipe/helper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

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

type operation func(ctx context.Context) error

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: helper.GetConfig("DB_USERNAME"),
		DB_PASSWORD: helper.GetConfig("DB_PASSWORD"),
		DB_HOST:     helper.GetConfig("DB_HOST"),
		DB_PORT:     helper.GetConfig("DB_PORT"),
		DB_NAME:     helper.GetConfig("DB_NAME"),
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

	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"database": func(ctx context.Context) error {
			return _dbDriver.CloseDB(db)
		},
		"http-server": func(ctx context.Context) error {
			return e.Shutdown(context.Background())
		},
	})

	<-wait
}

// gracefulShutdown performs application shut down gracefully.
func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		// add any other syscalls that you want to be notified with
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("shutting down")

		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		// Do the operations asynchronously to save time
		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("%s: clean up failed: %s", innerKey, err.Error())
					return
				}

				log.Printf("%s was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}
