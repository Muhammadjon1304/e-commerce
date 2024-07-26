package models

type User struct {
	ID            string `json:"id"`
	Username      string `json:"username" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	Password_hash string `json:"password_hash" binding:"required,min=2,max=10"`
	Role          string `json:"role" binding:"required"`
}
