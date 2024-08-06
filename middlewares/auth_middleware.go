package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/muhammadjon1304/e-commerce/utils"
	"log"
	"net/http"
	"os"
)

func AuthMiddleware() gin.HandlerFunc {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	return func(c *gin.Context) {
		// Retrieve the token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		// Define the claims structure
		claims := &utils.Claims{}

		// Parse the JWT token with claims
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Retrieve the JWT secret from the .env file
			jwtSecret := os.Getenv("JWT_SECRET")
			if jwtSecret == "" {
				return nil, fmt.Errorf("JWT secret not found in environment variables")
			}
			return []byte(jwtSecret), nil
		})

		// Check for errors in token parsing or validity
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set the username in the context for further handlers
		c.Set("username", claims.Username)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	if err := godotenv.Load(); err != nil {

		log.Fatalf("Error loading .env file")

	}
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}
		claims := &utils.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Retrieve the JWT secret from the .env file
			jwtSecret := os.Getenv("JWT_SECRET")
			if jwtSecret == "" {
				return nil, fmt.Errorf("JWT secret not found in environment variables")
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to access this resource"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
