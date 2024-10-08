package models

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       float64 `json:"stock"`
	Image_url   string  `json:"image_url"`
}

type ProductURI struct {
	ID uint `uri:"id" binding:"required,number"`
}

type ProductURICategory struct {
	ID uint `uri:"product_id" binding:"required,number"`
}

type CategoryURIProduct struct {
	ID uint `uri:"category_id" binding:"required,number"`
}
