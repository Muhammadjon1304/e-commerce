package repositories

import (
	"database/sql"
	"github.com/muhammadjon1304/e-commerce/models"
	"log"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return OrderRepository{DB: db}
}

func (o *OrderRepository) AddOrder(id uint, order models.OrderPost) bool {
	total_amount := 0
	stmt, err := o.DB.Prepare("INSERT INTO orders(user_id,total_amount,status) VALUES ($1,$2,$3)")
	if err != nil {
		log.Fatal(err)
		return false
	}
	_, err = stmt.Exec(id, total_amount, order.Status)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (c *OrderRepository) GetAllOrders(user uint) []models.Order {
	query, err := c.DB.Query("SELECT id,user_id,total_amount,status FROM orders WHERE user_id=$1", user)
	if err != nil {
		log.Fatal(err)
	}
	var Orders []models.Order

	if query != nil {
		for query.Next() {
			var (
				id           uint
				id_user      uint
				total_amount float64
				status       string
			)
			err := query.Scan(&id, &id_user, &total_amount, &status)
			if err != nil {
				log.Fatal(err)
			}
			order := models.Order{id, id_user, total_amount, status}
			Orders = append(Orders, order)
		}
	}
	return Orders
}

func (o *OrderRepository) GetOrderByID(user_id, id uint) models.Order {
	query, err := o.DB.Query("SELECT id,user_id,total_amount,status FROM orders WHERE id=$1 AND user_id=$2", id, user_id)

	if err != nil {
		log.Fatal(err)
		return models.Order{}
	}
	var order models.Order

	if query != nil {
		for query.Next() {
			var (
				id           uint
				user_id      uint
				total_amount float64
				status       string
			)
			err := query.Scan(&id, &user_id, &total_amount, &status)
			if err != nil {
				log.Fatal(err)
			}
			order = models.Order{id, user_id, total_amount, status}
		}
	}
	return order
}

func (o *OrderRepository) AddOrderItems(items models.Order_item) bool {
	stmt, err := o.DB.Prepare("INSERT INTO order_items VALUES (user_id,total_amount,status)")
	if err != nil {
		log.Fatal()
		return false
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(items.Order_id, items.Product_id, items.Quantity, items.Price)
		return false
	}
	return true
}

func (o *OrderRepository) GetOrderItems(order_id uint) []models.Order_item {
	query, err := o.DB.Query("SELECT id,order_id,product_id,quantity,price FROM order_items WHERE order_id=$1", order_id)
	if err != nil {
		log.Fatal(err)
	}
	var order_items []models.Order_item

	if query != nil {
		for query.Next() {
			var (
				id         uint
				order_id   uint
				product_id uint
				quantity   float64
				price      float64
			)
			err := query.Scan(&id, &order_id, &product_id, &quantity, &price)
			if err != nil {
				log.Fatal(err)
			}
			order_item := models.Order_item{id, order_id, product_id, quantity, price}
			order_items = append(order_items, order_item)
		}
	}
	return order_items
}

func (u *OrderRepository) GetUserIDByUsername(username string) uint {
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

func (o *OrderRepository) GetQuantity(orderID uint) float64 {
	query, err := o.DB.Query("SELECT total_amount FROM orders WHERE id=$1", orderID)
	if query != nil {
		log.Fatal(err)
	}
	var total_amount float64
	if query != nil {
		for query.Next() {
			var (
				amount float64
			)
			err := query.Scan(&amount)
			if err != nil {
				log.Fatal(err)
			}
			total_amount = amount
		}
	}
	return total_amount
}

func (o *OrderRepository) SetQuantity(amount float64, orderID uint) bool {
	quantity := o.GetQuantity(orderID)
	quantity += amount
	_, err := o.DB.Exec("UPDATE orders SET total_amount=$1", quantity)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
