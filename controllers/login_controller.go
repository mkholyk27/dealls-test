package controllers

import (
	"dating-app/models"
	"dating-app/repositories"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

var once sync.Once

type LoginController struct {
	repo *repositories.LoginRepository
}

var loginController *LoginController

func NewLoginController(repo *repositories.LoginRepository) *LoginController {
	once.Do(func() {
		loginController = &LoginController{
			repo: repo,
		}
	})

	return loginController
}

func (cntrl *LoginController) Login(ctx echo.Context) error {
	username, password, ok := ctx.Request().BasicAuth()
	if !ok {
		return ctx.JSON(http.StatusOK, models.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  false,
			Message: "Invalid username or password",
		})
	}

	result, err := cntrl.repo.Login(ctx, username, password)
	if err != nil {
		return ctx.JSON(http.StatusOK, models.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  false,
			Message: err.Error(),
		})
	}

	token := make(map[string]string)
	json.Unmarshal([]byte(result), &token)

	return ctx.JSON(http.StatusOK, models.WebResponse{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Login was successfull",
		Data:    token,
	})
}
