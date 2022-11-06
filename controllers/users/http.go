package users

import (
	"echo-recipe/app/middlewares"
	"echo-recipe/businesses/users"
	"echo-recipe/controllers/users/request"
	"echo-recipe/controllers/users/response"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authUseCase users.Usecase
}

func NewAuthController(authUC users.Usecase) *AuthController {
	return &AuthController{
		authUseCase: authUC,
	}
}

func (ctrl *AuthController) Register(c echo.Context) error {
	userInput := request.User{}

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	if err := c.Validate(&userInput); err != nil {
		return err
	}

	user := ctrl.authUseCase.Register(userInput.ToDomain())

	return c.JSON(http.StatusCreated, response.FromDomainRegister(user))
}

func (ctrl *AuthController) Login(c echo.Context) error {
	userInput := request.UserLogin{}

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	if err := c.Validate(&userInput); err != nil {
		return err
	}

	user := ctrl.authUseCase.Login(userInput.ToDomain().Email, userInput.Password)

	if user.Email == "" && user.Password == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid email or password",
		})
	}

	return c.JSON(http.StatusOK, response.FromDomain(user))

}

func (ctrl *AuthController) Logout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)

	isListed := middlewares.CheckToken(user.Raw)

	if !isListed {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid token",
		})
	}

	middlewares.Logout(user.Raw)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "logout success",
	})
}
