package repositories

import (
	"dating-app/helpers"
	"dating-app/models"
	"sync"

	"github.com/labstack/echo/v4"
)

var initSignup sync.Once

type SignupRepository struct {
}

var signupRepository *SignupRepository

func NewSignupRepository() *SignupRepository {
	initSignup.Do(func() {
		signupRepository = &SignupRepository{}
	})

	return signupRepository
}

func (repo *SignupRepository) Register(ctx echo.Context, data *models.User) error {
	err := helpers.InsertUserData(data)
	if err != nil {
		return err
	}

	return nil
}
