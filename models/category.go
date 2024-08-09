package models

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name" binding:"required"`
}

type CategoryURI struct {
	ID uint `uri:"id" binding:"required,number"`
}
