package models

type Order_item struct {
	ID         string `json:"id"`
	Order_id   string `json:"order_id"`
	Product_id string `json:"product_id"`
	Quantity   uint   `json:"quantity"`
	Price      uint   `json:"price"`
}
