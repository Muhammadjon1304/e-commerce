package models

type User struct {
	ID            uint   `json:"id"`
	Username      string `json:"username" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	Password_hash string `json:"password_hash" binding:"required,min=6,max=20"`
	Role          string `json:"role" binding:"required"`
}

type LoginUser struct {
	Username      string `json:"username" binding:"required"`
	Password_hash string `json:"password_hash" binding:"required,min=4,max=20"`
}
