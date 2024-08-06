package models

type CartItem struct {
	Id         uint    `json:"id"`
	Cart_id    uint    `json:"cart_id"`
	Product_id uint    `json:"product_id"`
	Quantity   float64 `json:"quantity"`
}

type PostCartItem struct {
	Product_id uint    `json:"product_id"`
	Quantity   float64 `json:"quantity"`
}

type CartItemURI struct {
	ID uint `uri:"id" binding:"required,number"`
}
