package middlewares

import (
	"echo-recipe/helper"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// AuthorizeJWT mengecek apakah tokennya valid atau tidak
func AuthorizeJWT(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		autheader := c.Request().Header.Get("Authorization")

		if !strings.Contains(autheader, "Bearer") {
			response := helper.BuildErrorResponse("failed to process request", "No token found", nil)
			return c.JSON(http.StatusUnauthorized, response)
		}

		jwtString := strings.Split(autheader, "Bearer ")[1]

		token, err := jwt.Parse(jwtString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("token not valid", err.Error(), nil)
			c.JSON(http.StatusUnauthorized, response)
		}
		return next(c)
	}
}
