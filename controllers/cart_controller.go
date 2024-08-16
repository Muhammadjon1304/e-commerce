package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/muhammadjon1304/e-commerce/models"
	"github.com/muhammadjon1304/e-commerce/repositories"
	"github.com/muhammadjon1304/e-commerce/status"
	"github.com/muhammadjon1304/e-commerce/views"
)

type CartController struct {
	DB *sql.DB
}

func NewCartController(db *sql.DB) CartController {
	return CartController{DB: db}
}

// CreateCart godoc
// @Summary Create a new cart for a user
// @Description Creates a new shopping cart for the authenticated user if one doesn't already exist
// @Tags Cart
// @Produce  json
// @Success 200 {object} R
// @Failure 400 {object} R
// @Failure 500 {object} R
// @Security ApiKeyAuth
// @Router /cart [post]
func (cc *CartController) CreateCart(c *gin.Context) {
	db := cc.DB
	username, exist := c.Get("username")
	if exist {
		repository := repositories.NewCartRepository(db)
		userID := repository.GetUserIDByUsername(username.(string))
		cartID, err := repository.CheckExistingCart(userID)
		if cartID != 0 && err != nil {
			c.JSON(500, views.ErrView(status.ErrorCodeDB, "Cart for this user already created"))
			return
		}
		cart := repository.CreateCart(userID)
		c.JSON(200, views.View(cart))
		return
	}
}

// GetCart godoc
// @Summary Retrieve the current cart for a user
// @Description Retrieves the shopping cart of the authenticated user if it exists
// @Tags Cart
// @Produce  json
// @Success 200 {object} R
// @Failure 400 {object} R
// @Failure 404 {object} R
// @Failure 500 {object} R
// @Security ApiKeyAuth
// @Router /cart [get]
func (cc *CartController) GetCart(c *gin.Context) {
	db := cc.DB
	username, exist := c.Get("username")
	if exist {
		repository := repositories.NewCartRepository(db)
		userID := repository.GetUserIDByUsername(username.(string))
		cartID, err := repository.CheckExistingCart(userID)
		if (cartID != 0) && (err == nil) {
			cart := repository.GetCart(userID)
			c.JSON(200, views.View(cart))
			return
		}
	}
}

// AddItemToCart godoc
// @Summary Add an item to the user's cart
// @Description Adds a specified item to the authenticated user's cart if the cart exists
// @Tags Cart
// @Accept  json
// @Produce  json
// @Param   cartItem  body  models.PostCartItem  true  "Cart Item JSON"
// @Success 200 {object} R
// @Failure 400 {object} R
// @Failure 404 {object} R
// @Failure 500 {object} R
// @Security ApiKeyAuth
// @Router /cart/items [post]
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
					ctx.JSON(200, views.View(nil))
					return
				} else {
					ctx.JSON(500, views.ErrView(status.ErrorCodeDB, "Item is not inserted to cart"))
					return
				}
			}
		}
	}
}

// UpdateCartItem godoc
// @Summary Update an item in the user's cart
// @Description Updates a specific item in the authenticated user's cart by item ID
// @Tags Cart
// @Accept  json
// @Produce  json
// @Param   itemID   path  int  true  "Item ID"
// @Param   cartItem body  models.PostCartItem  true  "Updated Cart Item JSON"
// @Success 200 {object} R
// @Failure 400 {object} R
// @Failure 401 {object} R
// @Failure 404 {object} R
// @Failure 500 {object} R
// @Security ApiKeyAuth
// @Router /cart/items/{itemID} [put]
func (c *CartController) UpdateCartItem(ctx *gin.Context) {
	db := c.DB
	username, exist := ctx.Get("username")
	if !exist {
		ctx.JSON(401, views.ErrView(status.ErrorCodeDB, "Username does not exist"))
		return
	}

	repository := repositories.NewCartRepository(db)

	var cartItem models.PostCartItem
	if err := ctx.ShouldBindJSON(&cartItem); err != nil {
		ctx.JSON(400, views.ErrView(status.ErrorCodeValidation, "Invalid payload"))
		return
	}

	var itemID models.CartItemURI
	if err := ctx.ShouldBindUri(&itemID); err != nil {
		ctx.JSON(400, views.ErrView(status.ErrorCodeValidation, "Item ID is not given"))
		return
	}

	userID := repository.GetUserIDByUsername(username.(string))
	if userID == 0 {
		ctx.JSON(404, views.ErrView(status.ErrorCodeDB, "User does not exist"))
		return
	}

	cartID, err := repository.CheckExistingCart(userID)
	if err != nil {
		ctx.JSON(500, views.ErrView(status.ErrorCodeDB, "Error checking card"))
		return
	}

	if cartID == 0 {
		ctx.JSON(404, views.ErrView(status.ErrorCodeDB, "Cart is not found"))
		return
	}

	updatedCartItem := repository.UpdateCartItem(cartID, itemID.ID, cartItem)
	if (updatedCartItem == models.CartItem{}) {
		ctx.JSON(500, views.ErrView(status.ErrorCodeDB, "Item is not updated"))
		return
	}

	ctx.JSON(200, views.View(nil))
}

// DeleteCartItem godoc
// @Summary Delete an item from the user's cart
// @Description Deletes a specific item from the authenticated user's cart by item ID
// @Tags Cart
// @Accept  json
// @Produce  json
// @Param   itemID   path  int  true  "Item ID"
// @Success 200 {object} R
// @Failure 400 {object} R
// @Failure 401 {object} R
// @Failure 404 {object} R
// @Failure 500 {object} R
// @Security ApiKeyAuth
// @Router /cart/items/{itemID} [delete]
func (c *CartController) DeleteCartItem(ctx *gin.Context) {
	db := c.DB
	username, exist := ctx.Get("username")

	if !exist {
		ctx.JSON(401, views.ErrView(status.ErrorCodeDB, "User is not found"))
		return
	}

	repository := repositories.NewCartRepository(db)

	var ItemID models.CartItemURI

	if err := ctx.ShouldBindUri(&ItemID); err != nil {
		ctx.JSON(400, views.ErrView(status.ErrorCodeValidation, "Item ID is not given"))
		return
	}

	userID := repository.GetUserIDByUsername(username.(string))
	if userID == 0 {
		ctx.JSON(404, views.ErrView(status.ErrorCodeDB, "User is not found"))
		return
	}

	cartID, err := repository.CheckExistingCart(userID)
	if err != nil {
		ctx.JSON(500, views.ErrView(status.ErrorCodeDB, err.Error()))
		return
	}

	if cartID == 0 {
		ctx.JSON(404, views.ErrView(status.ErrorCodeDB, "Cart is not found"))
		return
	}

	deleted := repository.DeleteCartItem(ItemID.ID, cartID)
	if !deleted {
		ctx.JSON(404, views.ErrView(status.ErrorCodeDB, "Item not deleted"))
		return
	}
	ctx.JSON(200, views.View(nil))
	return
}
