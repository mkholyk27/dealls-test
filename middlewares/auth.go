package middlewares

import (
	"dating-app/models"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("17p@ss27P@ss29Pass28pass!")

func MiddlewareJWTAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Path() == "/login" || ctx.Path() == "/register" {
			return next(ctx)
		}

		authorizationHeader := ctx.Request().Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			return ctx.JSON(http.StatusOK, models.WebResponse{
				Code:    http.StatusBadRequest,
				Status:  false,
				Message: "Invalid Token",
			})
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("signing method invalid")
			}

			return JWT_SIGNATURE_KEY, nil
		})
		if err != nil {
			return ctx.JSON(http.StatusOK, models.WebResponse{
				Code:    http.StatusBadRequest,
				Status:  false,
				Message: err.Error(),
			})
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return ctx.JSON(http.StatusOK, models.WebResponse{
				Code:    http.StatusBadRequest,
				Status:  false,
				Message: "Invalid Token",
			})
		}

		ctx.Set("userInfo", claims)
		return next(ctx)
	}
}
