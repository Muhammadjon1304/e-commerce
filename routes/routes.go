package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadjon1304/e-commerce/controllers"
	"github.com/muhammadjon1304/e-commerce/middlewares"
)

func UserRoutes(router *gin.Engine, ctrl *controllers.UserController) {
	router.POST("/users/register", ctrl.Register)
	router.POST("/users/login", ctrl.Login)
	router.GET("/users/profile", middlewares.AuthMiddleware(), ctrl.GetProfile)
}

func ProductRoutes(router *gin.Engine, ctrl *controllers.ProductController) {
	router.GET("/products", ctrl.GetAllProducts)
	router.GET("/products/:id", ctrl.GetProductByID)
	router.POST("/products", middlewares.AdminMiddleware(), ctrl.AddProduct)
	router.PUT("/products/:id", middlewares.AdminMiddleware(), ctrl.UpdateProduct)
	router.DELETE("/products/:id", middlewares.AdminMiddleware(), ctrl.DeleteProduct)
	router.POST("/products/:product_id/categories/:id", middlewares.AdminMiddleware(), ctrl.AddCategory)
	router.DELETE("/products/:id/categories/:id", middlewares.AdminMiddleware(), ctrl.DeleteCategory)
}

func CategoryRoutes(router *gin.Engine, ctrl *controllers.CategoryController) {
	router.GET("/categories", ctrl.GetAllCategory)
	router.POST("/categories", middlewares.AdminMiddleware(), ctrl.AddCategory)
	router.DELETE("/categories/:id", middlewares.AdminMiddleware(), ctrl.DeleteCategory)
}

func OrderRoutes(router *gin.Engine, ctrl *controllers.OrderController) {
	router.POST("/orders", middlewares.AuthMiddleware(), ctrl.AddOrder)
	router.GET("/orders", middlewares.AuthMiddleware(), ctrl.GetAllOrders)
	router.GET("/orders/:id", middlewares.AuthMiddleware(), ctrl.GetOrderDetails)
	router.POST("/orders/:id/items", middlewares.AuthMiddleware(), ctrl.AddItemToOrder)
	router.GET("/orders/:id/items", middlewares.AuthMiddleware(), ctrl.GetAllOrderItems)
}

func CartRoutes(router *gin.Engine, ctrl *controllers.CartController) {
	router.GET("/cart", middlewares.AuthMiddleware(), ctrl.GetCart)
	router.POST("/cart", middlewares.AuthMiddleware(), ctrl.CreateCart)
	router.POST("/cart/items", middlewares.AuthMiddleware(), ctrl.AddItemToCart)
	router.PUT("/cart/items/:id", middlewares.AuthMiddleware(), ctrl.UpdateCartItem)
	router.DELETE("/cart/items/:id", middlewares.AuthMiddleware(), ctrl.DeleteCartItem)
}
