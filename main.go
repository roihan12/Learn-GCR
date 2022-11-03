package main

import (
	"echo-recipe/helper"
	"echo-recipe/middlewares"
	"echo-recipe/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {

	server := echo.New()

	server.Validator = &helper.CustomValidator{Validator: validator.New()}

	middlewares.LogMiddleware(server)

	routes.SetupRoute(server)

	server.Logger.Fatal(server.Start(":1323"))
}
