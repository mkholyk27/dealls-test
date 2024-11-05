package repositories

import (
	"dating-app/helpers"
	"errors"
	"os"
	"path/filepath"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/gojsonq/v2"
)

var once sync.Once

type LoginRepository struct {
}

var loginRepository *LoginRepository

func NewLoginRepository() *LoginRepository {
	once.Do(func() {
		loginRepository = &LoginRepository{}
	})

	return loginRepository
}

func (repo *LoginRepository) Login(ctx echo.Context, username, password string) (string, error) {
	ok, userInfo := authenticateUser(username, password)
	if !ok {
		return "", errors.New("invalid username or password")
	}
	auth := new(helpers.Auth)

	return auth.GenerateJWT(userInfo)
}

func authenticateUser(username, password string) (bool, map[string]interface{}) {
	basePath, _ := os.Getwd()
	dbPath := filepath.Join(basePath+"/db/", "users.json")
	jq := gojsonq.New().File(dbPath)

	res := jq.From("users").Where("username", "=", username).Where("password", "=", password).First()

	if res != nil {
		resM := res.(map[string]interface{})
		delete(resM, "password")
		return true, resM
	}

	return false, nil
}
