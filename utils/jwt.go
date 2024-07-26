package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/muhammadjon1304/e-commerce/models"
	"log"
	"os"
	"time"
)

func GenerateJWT(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.Username,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}
