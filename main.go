package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/muhammadjon1304/e-commerce/config"
	"github.com/muhammadjon1304/e-commerce/controllers"
	_ "github.com/muhammadjon1304/e-commerce/docs"
	"github.com/muhammadjon1304/e-commerce/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title E-Commerce API
// @version 1.0
// @description This is a sample server for an e-commerce application.
func main() {
	_ = godotenv.Load()
	db := config.ConnectDB()
	router := gin.Default()
	UserController := controllers.NewUserController(db)
	ProductController := controllers.NewProductController(db)
	CategoryController := controllers.NewCategoryController(db)
	OrderController := controllers.NewOrderController(db)
	CartController := controllers.NewCartController(db)
	routes.UserRoutes(router, &UserController)
	routes.ProductRoutes(router, &ProductController)
	routes.CategoryRoutes(router, &CategoryController)
	routes.OrderRoutes(router, &OrderController)
	routes.CartRoutes(router, &CartController)

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")

}
