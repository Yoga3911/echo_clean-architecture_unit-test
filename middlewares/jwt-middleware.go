package middlewares

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func CreateJWTToken(id uint, name string) (string, error) {
	claims := jwt.MapClaims{
		"id": id,
		"name": name,
		"exp": time.Now().Add(10 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256 ,claims)

	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}