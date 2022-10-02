package middlewares

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/mock"
)

type JWTSMock interface {
	CreateJWTToken(id uint, name string) (string, error)
}

type IjwtSMock struct {
	issuer    string
	secretKey string
	Mock      mock.Mock
}

func NewJWTSMock(Mock mock.Mock) JWTSMock {
	err := godotenv.Load()
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			log.Println(err)
		}

		log.Println(err)
	}
	return &IjwtSMock{
		issuer:    "qwerty",
		secretKey: os.Getenv("JWT_KEY"),
		Mock:      Mock,
	}
}

func (j *IjwtSMock) CreateJWTToken(id uint, name string) (string, error) {
	args := j.Mock.Called()
	if args.Get(0) == nil {
		return "", args.Get(1).(error)
	}

	claims := jwt.MapClaims{
		"id":   id,
		"name": name,
		"exp":  time.Now().Add(10 * time.Minute).Unix(),
		"iss":  j.issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.secretKey))
	return t, err
}
