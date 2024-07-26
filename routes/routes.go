package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadjon1304/e-commerce/controllers"
)

func AllRoutes(router *gin.Engine, ctrl *controllers.UserController) {
	router.POST("users/register", ctrl.Register)
}
