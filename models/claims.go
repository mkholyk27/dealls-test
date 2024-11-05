package models

import jwt "github.com/golang-jwt/jwt/v4"

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Group    string `json:"Group"`
	IsMember int    `json:"IsMember"`
}
