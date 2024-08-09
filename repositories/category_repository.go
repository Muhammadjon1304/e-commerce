package repositories

import (
	"database/sql"
	"errors"
	"github.com/muhammadjon1304/e-commerce/models"
	"log"
)

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return CategoryRepository{DB: db}
}

func (c *CategoryRepository) GetAllCategories() []models.Category {
	query, err := c.DB.Query("SELECT id,name FROM categories")
	if err != nil {
		log.Fatal(err)
	}
	var Categories []models.Category

	if query != nil {
		for query.Next() {
			var (
				id   uint
				name string
			)
			err := query.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			category := models.Category{id, name}
			Categories = append(Categories, category)
		}
	}
	return Categories
}

func (c *CategoryRepository) AddCategory(category models.Category) bool {
	stmt, err := c.DB.Prepare("INSERT INTO categories(name) VALUES ($1)")
	if err != nil {
		log.Fatal()
		return false
	}
	_, err = stmt.Exec(category.Name)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (c *CategoryRepository) DeleteCategory(id uint) bool {
	_, err := c.DB.Query("DELETE FROM categories WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (c *CategoryRepository) AddCategoryToProduct(product_ID uint, category_id uint) bool {
	stmt, err := c.DB.Prepare("INSERT INTO product_categories(product_id,category_id) values($1,$2) ")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(product_ID, category_id)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

var db *sql.DB

func (c *CategoryRepository) CheckCategoryProductRelationship(productId, categoryId uint) (bool, error) {
	if c.DB == nil {
		return false, errors.New("database connection is not initialized")
	}

	var dummyID uint
	query := `SELECT product_id FROM product_categories WHERE product_id = $1 AND category_id = $2`

	err := c.DB.QueryRow(query, productId, categoryId).Scan(&dummyID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		log.Printf("Error checking category-product relationship: %v", err)
		return false, err
	}

	return true, nil
}

func (c *CategoryRepository) DeleteCategoryProductRelationship(productId, categoryId uint) bool {
	_, err := c.DB.Query("DELETE FROM product_categories WHERE product_id = $1 AND category_id = $2", productId, categoryId)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (u *CategoryRepository) GetUserIDByUsername(username string) uint {
	query, err := u.DB.Query("SELECT id FROM users WHERE username=$1", username)

	if err != nil {
		log.Fatal(err)
		return 0
	}
	var id uint

	if query != nil {
		for query.Next() {
			var (
				user_id uint
			)
			err := query.Scan(&user_id)
			if err != nil {
				log.Fatal(err)
			}
			id = user_id
		}
	}
	return id
}
