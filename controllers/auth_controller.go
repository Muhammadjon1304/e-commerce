package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/muhammadjon1304/e-commerce/models"
	"github.com/muhammadjon1304/e-commerce/repositories"
	"github.com/muhammadjon1304/e-commerce/utils"
	"golang.org/x/crypto/bcrypt"
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

func (u *UserController) Login(c *gin.Context) {
	var user models.LoginUser
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user_db models.User
	repository := repositories.NewUserRepository(u.DB)
	user_db = repository.GetUserByUsername(user.Username)

	if err := bcrypt.CompareHashAndPassword([]byte(user_db.Password_hash), []byte(user.Password_hash)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	token := utils.GenerateJWT(user_db)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (u *UserController) GetProfile(c *gin.Context) {
	username, exist := c.Get("username")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get username"})
		return
	}
	dbuser := repositories.NewUserRepository(u.DB)
	user := dbuser.GetUserByUsernameForUser(username.(string))

	c.JSON(http.StatusOK, gin.H{"user": user})
}
