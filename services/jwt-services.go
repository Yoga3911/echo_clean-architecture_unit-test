package services

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateJWTToken(id uint, name string) (string, error) {
	claims := jwt.MapClaims{
		"id": id,
		"name": name,
		"exp": time.Now().Add(10 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256 ,claims)

	return token.SignedString([]byte(""))
}