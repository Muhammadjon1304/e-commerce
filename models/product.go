package models

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       float64 `json:"stock"`
	Image_url   string  `json:"image_url"`
}
