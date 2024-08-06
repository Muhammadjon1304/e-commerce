package models

type Order struct {
	ID           uint    `json:"id"`
	User_id      uint    `json:"user"`
	Total_amount float64 `json:"total_amount"`
	Status       string  `json:"status"`
}

type OrderPost struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
}

type OrderURI struct {
	ID uint `uri:"id" binding:"required,number"`
}
