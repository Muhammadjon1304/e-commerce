package repositories

import (
	"database/sql"
	"github.com/muhammadjon1304/e-commerce/models"
	"log"
)

type ProductRepo struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) ProductRepo {
	return ProductRepo{DB: db}
}

func (p *ProductRepo) AddProduct(product models.Product) bool {
	query, err := p.DB.Prepare("INSERT INTO products(name,description,price,stock,image_url) values($1,$2,$3,$4,$5)")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer query.Close()
	_, err = query.Exec(product.Name, product.Description, product.Price, product.Stock, product.Image_url)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (p *ProductRepo) GetAllProducts() []models.Product {
	query, err := p.DB.Query("SELECT id,name,description,price,stock,image_url FROM products")
	if err != nil {
		log.Fatal(err)
	}
	var products []models.Product

	if query != nil {
		for query.Next() {
			var (
				id          uint
				name        string
				description string
				price       float64
				stock       float64
				image_url   string
			)
			err := query.Scan(&id, &name, &description, &price, &stock, &image_url)
			if err != nil {
				log.Fatal(err)
			}
			product := models.Product{id, name, description, price, stock, image_url}
			products = append(products, product)
		}
	}
	return products
}

func (p *ProductRepo) GetProductByID(id uint) models.Product {
	query, err := p.DB.Query("SELECT id,name,description,price,stock,image_url FROM products WHERE id=$1", id)

	if err != nil {
		log.Fatal(err)
		return models.Product{}
	}
	var product models.Product

	if query != nil {
		for query.Next() {
			var (
				id          uint
				name        string
				description string
				price       float64
				stock       float64
				image_url   string
			)
			err := query.Scan(&id, &name, &description, &price, &stock, &image_url)
			if err != nil {
				log.Fatal(err)
			}
			product = models.Product{id, name, description, price, stock, image_url}
		}
	}
	return product
}

func (p *ProductRepo) DeleteProduct(id uint) bool {
	_, err := p.DB.Exec("DELETE FROM products WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (p *ProductRepo) UpdateProduct(id uint, product models.Product) models.Product {
	_, err := p.DB.Exec("UPDATE products SET name=$1,description=$2,price=$3,stock=$4,image_url=$5", product.Name, product.Description, product.Price, product.Stock, product.Image_url)
	if err != nil {
		log.Fatal(err)
		return models.Product{}
	}
	return p.GetProductByID(id)
}
