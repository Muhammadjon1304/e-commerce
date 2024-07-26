package models

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name" binding:"reuired"`
}
