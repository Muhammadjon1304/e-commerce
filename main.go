package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/muhammadjon1304/e-commerce/config"
	"github.com/muhammadjon1304/e-commerce/controllers"
	"github.com/muhammadjon1304/e-commerce/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title E-Commerce API
// @version 1.0
// @description This is a sample server for an e-commerce application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

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
	router.Run(":8080")

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
