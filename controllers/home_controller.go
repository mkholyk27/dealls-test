package controllers

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type HomeController struct {
}

var homeController *HomeController

func NewHomeController() *HomeController {
	once.Do(func() {

		homeController = &HomeController{}
	})

	return homeController
}

func (cntrl *HomeController) Index(ctx echo.Context) error {
	userInfo := ctx.Get("userInfo").(jwt.MapClaims)
	message := fmt.Sprintf("Wellcome %s", userInfo["Username"])

	return ctx.String(http.StatusOK, message)
}
