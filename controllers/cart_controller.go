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
	if !exist {
		ctx.JSON(401, gin.H{"status": "fail", "message": "unauthorized"})
		return
	}

	repository := repositories.NewCartRepository(db)

	var cartItem models.PostCartItem
	if err := ctx.ShouldBindJSON(&cartItem); err != nil {
		ctx.JSON(400, gin.H{"status": "fail", "message": "invalid JSON"})
		return
	}

	var itemID models.CartItemURI
	if err := ctx.ShouldBindUri(&itemID); err != nil {
		ctx.JSON(400, gin.H{"status": "fail", "message": "invalid URI"})
		return
	}

	userID := repository.GetUserIDByUsername(username.(string))
	if userID == 0 {
		ctx.JSON(404, gin.H{"status": "fail", "message": "user not found"})
		return
	}

	cartID, err := repository.CheckExistingCart(userID)
	if err != nil {
		ctx.JSON(500, gin.H{"status": "fail", "message": "error checking cart"})
		return
	}

	if cartID == 0 {
		ctx.JSON(404, gin.H{"status": "fail", "message": "cart not found"})
		return
	}

	updatedCartItem := repository.UpdateCartItem(cartID, itemID.ID, cartItem)
	if (updatedCartItem == models.CartItem{}) {
		ctx.JSON(500, gin.H{"status": "fail", "message": "item not updated"})
		return
	}

	ctx.JSON(200, gin.H{"status": "success", "message": "item updated"})
}

func (c *CartController) DeleteCartItem(ctx *gin.Context) {
	db := c.DB
	username, exist := ctx.Get("username")

	if !exist {
		ctx.JSON(401, gin.H{"status": "fail", "message": "unauthorized"})
		return
	}

	repository := repositories.NewCartRepository(db)

	var ItemID models.CartItemURI

	if err := ctx.ShouldBindUri(&ItemID); err != nil {
		ctx.JSON(400, gin.H{"status": "fail", "message": "invalid uri"})
		return
	}

	userID := repository.GetUserIDByUsername(username.(string))
	if userID == 0 {
		ctx.JSON(404, gin.H{"status": "fail", "message": "user not found"})
		return
	}

	cartID, err := repository.CheckExistingCart(userID)
	if err != nil {
		ctx.JSON(500, gin.H{"status": "fail", "message": "error checking cart"})
		return
	}

	if cartID == 0 {
		ctx.JSON(404, gin.H{"status": "fail", "message": "cart not found"})
		return
	}

	deleted := repository.DeleteCartItem(ItemID.ID, cartID)
	if !deleted {
		ctx.JSON(404, gin.H{"status": "fail", "message": "cart not deleted"})
		return
	}
	ctx.JSON(200, gin.H{"status": "success", "message": "item deleted"})
	return
}
