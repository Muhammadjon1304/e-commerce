package models

type Order struct {
	ID           string `json:"id"`
	User_id      string `json:"user"`
	Total_amount uint   `json:"total_amount"`
	Status       string `json:"status"`
}
