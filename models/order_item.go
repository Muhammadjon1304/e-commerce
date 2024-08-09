package models

type Order_item struct {
	ID         uint    `json:"id"`
	Order_id   uint    `json:"order_id"`
	Product_id uint    `json:"product_id"`
	Quantity   float64 `json:"quantity"`
	Price      float64 `json:"price"`
}

type OrderItemURI struct {
	ID uint `uri:"id" binding:"required,number"`
}

type Order_itemPost struct {
	Quantity float64 `json:"quantity"`
	Price    float64 `json:"price"`
}
