package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadjon1304/e-commerce/controllers"
	"github.com/muhammadjon1304/e-commerce/middlewares"
)

func UserRoutes(router *gin.Engine, ctrl *controllers.UserController) {
	router.POST("users/register", ctrl.Register)
	router.POST("users/login", ctrl.Login)
	router.GET("users/profile", middlewares.AuthMiddleware(), ctrl.GetProfile)
}

func ProductRoutes(router *gin.Engine, ctrl *controllers.ProductController) {
	router.GET("products", ctrl.GetAllProducts)
	router.GET("/products/:id", ctrl.GetProductByID)
	router.POST("products", middlewares.AdminMiddleware(), ctrl.AddProduct)
	router.PUT("/products/:id", middlewares.AdminMiddleware(), ctrl.UpdateProduct)
	router.DELETE("/products/:id", middlewares.AdminMiddleware(), ctrl.DeleteProduct)
}
