package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadjon1304/e-commerce/controllers"
	"github.com/muhammadjon1304/e-commerce/middlewares"
)

func AllRoutes(router *gin.Engine, ctrl *controllers.UserController) {
	router.POST("users/register", ctrl.Register)
	router.POST("users/login", ctrl.Login)
	router.GET("users/profile", middlewares.AuthMiddleware(), ctrl.GetProfile)
}
