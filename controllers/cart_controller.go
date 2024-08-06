package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/muhammadjon1304/e-commerce/models"
	"github.com/muhammadjon1304/e-commerce/repositories"
)

type CartController struct {
	DB *sql.DB
}

func NewCartController(db *sql.DB) CartController {
	return CartController{DB: db}
}

func (cc *CartController) CreateCart(c *gin.Context) {
	db := cc.DB
	username, exist := c.Get("username")
	if exist {
		repository := repositories.NewCartRepository(db)
		userID := repository.GetUserIDByUsername(username.(string))
		cartID, err := repository.CheckExistingCart(userID)
		if cartID != 0 && err != nil {
			c.JSON(500, gin.H{"status": "fail", "msg": "cart not created"})
			return
		}
		cart := repository.CreateCart(userID)
		c.JSON(200, gin.H{"status": "success", "cart": cart})
		return
	}
}

func (cc *CartController) GetCart(c *gin.Context) {
	db := cc.DB
	username, exist := c.Get("username")
	if exist {
		repository := repositories.NewCartRepository(db)
		userID := repository.GetUserIDByUsername(username.(string))
		cartID, err := repository.CheckExistingCart(userID)
		if (cartID != 0) && (err == nil) {
			cart := repository.GetCart(userID)
			c.JSON(200, gin.H{"status": "success", "cart": cart})
			return
		}
	}
}

func (c *CartController) AddItemToCart(ctx *gin.Context) {
	db := c.DB
	username, exist := ctx.Get("username")
	fmt.Println(username, exist)
	if exist {
		repository := repositories.NewCartRepository(db)
		var cartItem models.PostCartItem
		err := ctx.BindJSON(&cartItem)
		fmt.Println(err)
		if err == nil {
			userID := repository.GetUserIDByUsername(username.(string))
			cartID, err := repository.CheckExistingCart(userID)
			if cartID != 0 && err == nil {
				inserted := repository.AddCartItems(cartID, cartItem)
				if inserted {
					ctx.JSON(200, gin.H{"status": "success", "message": "item inserted"})
					return
				} else {
					ctx.JSON(500, gin.H{"status": "fail", "message": "item not inserted"})
					return
				}
			}
		}
	}
}

func (c *CartController) UpdateCartItem(ctx *gin.Context) {
	db := c.DB
	username, exist := ctx.Get("username")
	if exist {
		repository := repositories.NewCartRepository(db)
		var cartItem models.PostCartItem
		err := ctx.ShouldBindJSON(&cartItem)
		fmt.Println(err)
		if err == nil {
			var itemID models.CartItemURI
			err = ctx.ShouldBindUri(itemID)
			if err == nil {
				userID := repository.GetUserIDByUsername(username.(string))
				cartID, err := repository.CheckExistingCart(userID)
				if cartID != 0 && err == nil {
					cartItem := repository.UpdateCartItem(cartID, itemID.ID, cartItem)
					if (cartItem != models.CartItem{}) {
						ctx.JSON(200, gin.H{"status": "success", "message": "item updated"})
						return
					} else {
						ctx.JSON(500, gin.H{"status": "fail", "message": "item not updated"})
						return
					}
				}
			}
		}
	}
}

func (c *CartController) DeleteCartItem(ctx *gin.Context) {
	db := c.DB
	username, exist := ctx.Get("username")
	if exist {
		repository := repositories.NewCartRepository(db)

		var itemID models.CartItemURI
		if err := ctx.ShouldBindUri(&itemID); err != nil {
			userID := repository.GetUserIDByUsername(username.(string))
			cartID, err := repository.CheckExistingCart(userID)
			if cartID != 0 && err != nil {
				deleted := repository.DeleteCartItem(itemID.ID, cartID)
				if deleted {
					ctx.JSON(200, gin.H{"status": "success", "message": "item deleted"})
					return
				} else {
					ctx.JSON(500, gin.H{"status": "fail", "message": "item not deleted"})
					return
				}
			}
		}
	}
}
