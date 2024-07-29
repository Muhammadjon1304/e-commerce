package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/muhammadjon1304/e-commerce/models"
	"log"
	"os"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(user models.User) string {
	if err := godotenv.Load(); err != nil {

		log.Fatalf("Error loading .env file")

	}
	expiration_time := time.Now().Add(time.Hour * 24)
	claims := Claims{
		Username:         user.Username,
		Role:             user.Role,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expiration_time)},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}
