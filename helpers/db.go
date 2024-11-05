package helpers

import (
	"dating-app/models"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/thedevsaddam/gojsonq/v2"
)

func InsertUserData(data *models.User) error {
	basePath, _ := os.Getwd()
	dbPath := filepath.Join(basePath+"/db/", "users.json")

	err := checkFile(dbPath)
	if err != nil {
		return err
	}

	userIsExist := checkUserExist(data.UserName, data.Email, dbPath)
	if userIsExist {
		return errors.New("username or email already exists")
	}

	jq := gojsonq.New().File(dbPath)

	file := jq.From("users").Get()
	if file == nil {
		return errors.New("empty data")
	}

	userJson, err := json.Marshal(file)
	if err != nil {
		return err
	}

	users := []models.User{}
	json.Unmarshal(userJson, &users)
	users = append(users, *data)

	newData := make(map[string]interface{})
	newData["users"] = users

	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(newData)
	if err != nil {
		return err
	}

	err = os.WriteFile(dbPath, dataBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkUserExist(username string, email string, dbPath string) bool {
	jq := gojsonq.New().File(dbPath)

	file := jq.From("users").Where("username", "=", username).OrWhere("email", "=", email).First()
	return file != nil
}
