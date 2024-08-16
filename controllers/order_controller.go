package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/muhammadjon1304/e-commerce/models"
	"github.com/muhammadjon1304/e-commerce/repositories"
	"github.com/muhammadjon1304/e-commerce/status"
	"github.com/muhammadjon1304/e-commerce/views"
	"net/http"
)

type OrderController struct {
	DB *sql.DB
}

func NewOrderController(db *sql.DB) OrderController {
	return OrderController{DB: db}
}

// AddOrder godoc
// @Summary Add a new order
// @Description Creates a new order for the authenticated user
// @Tags Order
// @Accept  json
// @Produce  json
// @Param   order  body  models.OrderPost  true  "Order JSON"
// @Success 200 {object} R
// @Failure 400 {object} R
// @Failure 500 {object} R
// @Router /orders [post]
func (o *OrderController) AddOrder(c *gin.Context) {
	db := o.DB
	username, exist := c.Get("username")

	if exist {
		var order models.OrderPost
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusInternalServerError, views.ErrView(status.ErrorCodeValidation,err.Error()))
			c.AbortWithStatus(500)
		}
		repository := repositories.NewOrderRepository(db)
		id := repository.GetUserIDByUsername(username.(string))
		fmt.Println(id, username)
		inserted := repository.AddOrder(id, order)
		if inserted {
			c.JSON(200, views.View(nil))
			return
		} else {
			c.JSON(500, views.ErrView(status.ErrorCodeDB,"Order is not added"))
			return
		}
	}
}

// GetAllOrders godoc
// @Summary Get all orders for the authenticated user
// @Description Retrieves all orders placed by the authenticated user
// @Tags Order
// @Produce  json
// @Success 200 {array} R
// @Failure 401 {object} R
// @Failure 500 {object} R
// @Router /orders [get]
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

// GetOrderDetails godoc
// @Summary Get order details
// @Description Retrieves details of a specific order by order ID for the authenticated user
// @Tags Order
// @Produce  json
// @Param   orderID  path  int  true  "Order ID"
// @Success 200 {object} R
// @Failure 400 {object} R
// @Failure 401 {object} R
// @Failure 404 {object} R
// @Failure 500 {object} R
// @Router /orders/{orderID} [get]
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
			c.JSON(200, views.View(order))
			return
		} else {
			c.JSON(http.StatusInternalServerError, views.ErrView(status.ErrorCodeDB,"Order can't obtained "))
		}
	}
}

// GetAllOrderItems godoc
// @Summary Get all items for a specific order
// @Description Retrieves all items for a specific order by order ID for the authenticated user
// @Tags Order
// @Produce  json
// @Param   orderID  path  int  true  "Order ID"
// @Success 200 {array} R
// @Failure 400 {object} R
// @Failure 401 {object} R
// @Failure 500 {object} R
// @Router /orders/{orderID}/items [get]
func (o *OrderController) GetAllOrderItems(c *gin.Context) {
	db := o.DB
	_, exist := c.Get("username")
	if exist {
		var OrderID models.OrderURI
		repository := repositories.NewOrderRepository(db)
		items := repository.GetOrderItems(OrderID.ID)
		c.JSON(200, views.View(items))
	}
}

// AddItemToOrder godoc
// @Summary Add an item to an order
// @Description Adds a new item to an existing order
// @Tags Orders
// @Accept  json
// @Produce  json
// @Param  orderID  path  int  true  "Order ID"
// @Param  item     body  models.Order_item  true  "Order Item JSON"
// @Success 200 {object} R
// @Failure 400 {object} R
// @Failure 401 {object} R
// @Failure 500 {object} R
// @Router /orders/{orderID}/items [post]
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
					c.JSON(200, views.View(nil))
					return
				} else {
					c.JSON(500,views.ErrView(status.ErrorCodeDB,"Item is not inserted"))
					return
				}
			}
		}
	}
}
