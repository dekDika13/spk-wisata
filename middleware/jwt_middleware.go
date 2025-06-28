package middleware

import (
	"backend/utils"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(userID uint, role uint, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["role"] = role
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}



func ClaimData(c echo.Context, field string) (interface{}, error) {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req := claims[field]

	return req, nil
}

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role, _ := ClaimData(c, "role")

		conv := fmt.Sprintf("%v", role)
		if conv != "1" {
			return c.JSON(http.StatusForbidden, utils.Response{
				Message: "you dont have permission to access",
				Code:    http.StatusForbidden,
			})
		}
		return next(c)
	}
}

func RoleJwt(requiredRole int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userToken := c.Get("user")
			if userToken == nil {
				return c.JSON(http.StatusUnauthorized, utils.Response{
					Message: "unauthorized: missing token",
					Code:    http.StatusUnauthorized,
				})
			}

			token, ok := userToken.(*jwt.Token)
			if !ok {
				return c.JSON(http.StatusUnauthorized, utils.Response{
					Message: "unauthorized: invalid token",
					Code:    http.StatusUnauthorized,
				})
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				return c.JSON(http.StatusUnauthorized, utils.Response{
					Message: "unauthorized: invalid claims",
					Code:    http.StatusUnauthorized,
				})
			}

			roleVal, ok := claims["role"].(float64)
			if !ok {
				return c.JSON(http.StatusForbidden, utils.Response{
					Message: "invalid role format",
					Code:    http.StatusForbidden,
				})
			}

			if int(roleVal) != requiredRole {
				return c.JSON(http.StatusForbidden, utils.Response{
					Message: "your role doesn't have permission to access this resource",
					Code:    http.StatusForbidden,
				})
			}

			return next(c)
		}
	}
}