package helpers

import (
	"dating-app/configs"
	"dating-app/models"
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY []byte

type Auth struct {
}

func (auth *Auth) GenerateJWT(userInfo map[string]interface{}) (string, error) {
	cfg := configs.New()
	JWT_SIGNATURE_KEY = []byte(cfg.Get("JWT_SIGNATURE_KEY"))

	claims := models.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.Get("APPLICATION_NAME"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(LOGIN_EXPIRATION_DURATION)),
		},
		Username: userInfo["username"].(string),
		Email:    userInfo["email"].(string),
		Group:    userInfo["group"].(string),
		IsMember: userInfo["isMember"].(int),
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}

	tokenString, _ := json.Marshal(map[string]interface{}{"token": signedToken})

	return string(tokenString), nil
}
