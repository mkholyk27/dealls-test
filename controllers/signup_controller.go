package controllers

import (
	"dating-app/models"
	"dating-app/repositories"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

var initSignup sync.Once

type SignupController struct {
	repo *repositories.SignupRepository
}

var signupController *SignupController

func NewSignupController(repo *repositories.SignupRepository) *SignupController {
	initSignup.Do(func() {
		signupController = &SignupController{
			repo: repo,
		}
	})

	return signupController
}

func (cntrl *SignupController) Register(ctx echo.Context) error {
	var request models.UserDTO
	valid := ctx.Bind(&request)
	if valid != nil {
		ctx.Logger().Panic(valid)
	}

	user := models.User{
		UserName: request.UserName,
		Password: request.Password,
		Email:    request.Email,
		Group:    "user",
		IsMember: 0,
	}

	err := cntrl.repo.Register(ctx, &user)
	if err != nil {
		return ctx.JSON(http.StatusOK, models.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  false,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.WebResponse{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Saving new user was successfull",
		Data:    request,
	})
}
