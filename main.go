package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadjon1304/e-commerce/config"
	"github.com/muhammadjon1304/e-commerce/controllers"
	"github.com/muhammadjon1304/e-commerce/routes"
)

func main() {
	db := config.ConnectDB()
	router := gin.Default()
	controller := controllers.NewUserController(db)
	routes.AllRoutes(router, &controller)
	router.Run(":8080")

}
