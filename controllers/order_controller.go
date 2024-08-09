package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/muhammadjon1304/e-commerce/models"
	"github.com/muhammadjon1304/e-commerce/repositories"
	"net/http"
)

type OrderController struct {
	DB *sql.DB
}

func NewOrderController(db *sql.DB) OrderController {
	return OrderController{DB: db}
}

func (o *OrderController) AddOrder(c *gin.Context) {
	db := o.DB
	username, exist := c.Get("username")

	if exist {
		var order models.OrderPost
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.AbortWithStatus(500)
		}
		repository := repositories.NewOrderRepository(db)
		id := repository.GetUserIDByUsername(username.(string))
		fmt.Println(id, username)
		inserted := repository.AddOrder(id, order)
		if inserted {
			c.JSON(200, gin.H{"status": "success", "msg": "order created"})
			return
		} else {
			c.JSON(500, gin.H{"status": "fail", "msg": "order not created"})
			return
		}
	}
}

func (o *OrderController) GetAllOrders(c *gin.Context) {
	db := o.DB
	username, exist := c.Get("username")
	if exist {
		repository := repositories.NewOrderRepository(db)
		userID := repository.GetUserIDByUsername(username.(string))
		orders := repository.GetAllOrders(userID)
		c.JSON(200, gin.H{"orders": orders})
	}
}

func (o *OrderController) GetOrderDetails(c *gin.Context) {
	db := o.DB
	username, exist := c.Get("username")
	if exist {
		var orderID models.OrderURI
		if err := c.ShouldBindUri(&orderID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		fmt.Println(orderID, username)
		repository := repositories.NewOrderRepository(db)
		userID := repository.GetUserIDByUsername(username.(string))
		order := repository.GetOrderByID(userID, orderID.ID)
		if (order != models.Order{}) {
			c.JSON(200, gin.H{"status": "success", "order": order})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "error": "Can't get order"})
		}
	}
}

func (o *OrderController) GetAllOrderItems(c *gin.Context) {
	db := o.DB
	_, exist := c.Get("username")
	if exist {
		var OrderID models.OrderURI
		repository := repositories.NewOrderRepository(db)
		items := repository.GetOrderItems(OrderID.ID)
		c.JSON(200, gin.H{"order_items": items})
	}
}

func (o *OrderController) AddItemToOrder(c *gin.Context) {
	db := o.DB
	_, exist := c.Get("username")
	if exist {
		var OrderID models.OrderURI
		if err := c.ShouldBindUri(&OrderID); err != nil {
			var Item models.Order_item
			if err = c.ShouldBindJSON(&Item); err != nil {
				repository := repositories.NewOrderRepository(db)
				inserted := repository.AddOrderItems(Item)
				if inserted {
					quantity := Item.Quantity
					repository.SetQuantity(quantity, OrderID.ID)
					c.JSON(200, gin.H{"status": "success", "msg": "item inserted"})
					return
				} else {
					c.JSON(500, gin.H{"status": "fail", "msg": "item not inserted"})
					return
				}
			}
		}

	}

}
