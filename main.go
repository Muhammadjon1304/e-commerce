package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/muhammadjon1304/e-commerce/config"
	"github.com/muhammadjon1304/e-commerce/controllers"
	"github.com/muhammadjon1304/e-commerce/routes"
	"log"
	"os"
)

func main() {
	db := config.ConnectDB()
	router := gin.Default()
	controller := controllers.NewUserController(db)
	routes.AllRoutes(router, &controller)
	router.Run(":8080")

	if err := godotenv.Load(); err != nil {

		log.Fatalf("Error loading .env file")

	}
	fmt.Println(os.Getenv("JWT_SECRET"))
}
