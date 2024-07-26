package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/muhammadjon1304/e-commerce/models"
	"github.com/muhammadjon1304/e-commerce/repositories"
	"github.com/muhammadjon1304/e-commerce/utils"
	"net/http"
)

type UserController struct {
	DB *sql.DB
}

func NewUserController(db *sql.DB) UserController {
	return UserController{
		DB: db,
	}
}
func (u *UserController) Register(ctx *gin.Context) {
	db := u.DB
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}
	hashedPassword := utils.HashPassword(user.Password_hash)

	newUser := models.User{Username: user.Username, Email: user.Email, Password_hash: hashedPassword, Role: user.Role}

	repo := repositories.NewUserRepository(db)
	inserted := repo.SaveUser(newUser)

	if inserted {
		ctx.JSON(200, gin.H{"status": "success", "msg": "user created"})
		return
	} else {
		ctx.JSON(500, gin.H{"status": "fail", "msg": "user not created"})
	}
}

func Login() {

}
