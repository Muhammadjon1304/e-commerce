package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/muhammadjon1304/e-commerce/models"
	"github.com/muhammadjon1304/e-commerce/repositories"
	"github.com/muhammadjon1304/e-commerce/status"
	"github.com/muhammadjon1304/e-commerce/utils"
	"github.com/muhammadjon1304/e-commerce/views"
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

// Register godoc
// @Summary Register a new user
// @Description Register a new user with the input payload
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   user  body   models.User  true  "User JSON"
// @Success 200 {object} R
// @Failure 400 {object} R
// @Failure 500 {object} R
// @Router /users/register [post]
func (u *UserController) Register(ctx *gin.Context) {
	db := u.DB
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, views.ErrView(status.ErrorCodeValidation, err.Error()))
		return
	}
	hashedPassword := utils.HashPassword(user.Password_hash)

	newUser := models.User{Username: user.Username, Email: user.Email, Password_hash: hashedPassword, Role: user.Role}

	repo := repositories.NewUserRepository(db)
	inserted := repo.SaveUser(newUser)

	if inserted {
		ctx.JSON(200, views.View(nil))
		return
	} else {
		ctx.JSON(500, views.ErrView(status.ErrorCodeDB, "User is not inserted to database"))
	}
}

// Login godoc
// @Summary Log in a user
// @Description Authenticate a user with a username and password, returning a JWT if successful
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   loginUser  body  models.LoginUser  true  "Login User JSON"
// @Success 200 {object} R
// @Failure 400 {object} R
// @Failure 401 {object} R
// @Router /users/login [post]
func (u *UserController) Login(c *gin.Context) {
	var user models.LoginUser
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, views.ErrView(status.ErrorCodeValidation, "Invalid payload"))
		return
	}
	var user_db models.User
	repository := repositories.NewUserRepository(u.DB)
	user_db = repository.GetUserByUsername(user.Username)

	if err := bcrypt.CompareHashAndPassword([]byte(user_db.Password_hash), []byte(user.Password_hash)); err != nil {
		c.JSON(http.StatusUnauthorized, views.ErrView(status.ErrorCodeDB, "Password is not correct"))
		return
	}
	token := utils.GenerateJWT(user_db)

	c.JSON(http.StatusOK, views.View(token))
}

// GetProfile godoc
// @Summary Get user profile
// @Description Retrieve the profile information of the authenticated user
// @Tags Users
// @Produce  json
// @Success 200 {object} R
// @Failure 500 {object} R
// @Security ApiKeyAuth
// @Router /users/profile [get]
func (u *UserController) GetProfile(c *gin.Context) {
	username, exist := c.Get("username")
	if !exist {
		c.JSON(http.StatusInternalServerError, views.ErrView(status.ErrorCodeDB, "User is not exist"))
		return
	}
	dbuser := repositories.NewUserRepository(u.DB)
	user := dbuser.GetUserByUsernameForUser(username.(string))

	c.JSON(http.StatusOK, views.View(user))
}
