package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/muhammadjon1304/e-commerce/config"
	"github.com/muhammadjon1304/e-commerce/controllers"
	"github.com/muhammadjon1304/e-commerce/routes"
)

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
	routes.CartRoutes(router,&CartController)
	router.Run(":8080")
}
