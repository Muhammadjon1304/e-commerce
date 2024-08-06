package repositories

import (
	"database/sql"
	"github.com/muhammadjon1304/e-commerce/models"
	"log"
)

type CartRepository struct {
	DB *sql.DB
}

func NewCartRepository(db *sql.DB) CartRepository {
	return CartRepository{
		DB: db,
	}
}

func (c *CartRepository) CreateCart(UserID uint) bool {
	stmt, err := c.DB.Prepare("INSERT INTO carts(user_id) VALUES ($1)")
	if err != nil {
		log.Fatal()
		return false
	}
	_, err = stmt.Exec(UserID)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (c *CartRepository) GetCart(userID uint) models.Cart {
	query, err := c.DB.Query("SELECT id,user_id 	FROM carts WHERE user_id=$1", userID)
	if err != nil {
		log.Fatal(err)
		return models.Cart{}
	}
	defer query.Close()
	var cart models.Cart

	if query != nil {
		for query.Next() {
			var (
				id      uint
				user_id uint
			)
			err := query.Scan(&id, &user_id)
			if err != nil {
				log.Fatal(err)
			}
			cart = models.Cart{id, user_id}
		}
	}
	return cart
}

func (c *CartRepository) CheckExistingCart(userID uint) (uint, error) {
	var cartID uint
	err := c.DB.QueryRow("SELECT id FROM carts WHERE user_id = $1", userID).Scan(&cartID)
	if err == sql.ErrNoRows {
		return 0, nil // No existing cart
	}
	return cartID, err
}

func (c *CartRepository) AddCartItems(cartID uint, cart_item models.PostCartItem) bool {
	stmt, err := c.DB.Prepare("INSERT INTO cart_items(cart_id,product_id,quantity) VALUES ($1,$2,$3)")
	if err != nil {
		log.Fatal(err)
		return false
	}
	_, err = stmt.Exec(cartID, cart_item.Product_id, cart_item.Quantity)
	if err != nil {
		log.Fatal()
		return false
	}
	return true
}

func (c *CartRepository) DeleteCartItem(id, cart_id uint) bool {
	_, err := c.DB.Query("DELETE FROM cart_items WHERE id=$1 and cart_id=$2", id, cart_id)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (c *CartRepository) UpdateCartItem(cartID, id uint, cart_item models.PostCartItem) models.CartItem {
	_, err := c.DB.Exec("UPDATE cart_items SET product_id=$1,quantity=$2 WHERE cart_id=$3 AND id=4", cart_item.Product_id, cart_item.Quantity, cartID, id)
	if err != nil {
		log.Fatal(err)
		return models.CartItem{}
	}
	return c.GetCartItem(id)
}

func (c *CartRepository) GetCartItem(id uint) models.CartItem {
	query, err := c.DB.Query("SELECT id,cart_id,product_id,quantity FROM cart_items WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
		return models.CartItem{}
	}
	defer query.Close()
	var cart_item models.CartItem

	if query != nil {
		for query.Next() {
			var (
				id         uint
				cart_id    uint
				product_id uint
				quantity   float64
			)
			err := query.Scan(&id, &cart_id, &product_id, &quantity)
			if err != nil {
				log.Fatal(err)
			}
			cart_item = models.CartItem{id, cart_id, product_id, quantity}
		}
	}
	return cart_item
}

func (u *CartRepository) GetUserIDByUsername(username string) uint {
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
