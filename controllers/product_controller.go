package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/muhammadjon1304/e-commerce/models"
	"github.com/muhammadjon1304/e-commerce/repositories"
	"github.com/muhammadjon1304/e-commerce/status"
	"github.com/muhammadjon1304/e-commerce/views"
	"net/http"
)

type ProductController struct {
	DB *sql.DB
}

func NewProductController(db *sql.DB) ProductController {
	return ProductController{DB: db}
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Retrieves all available products from the database
// @Tags Product
// @Produce  json
// @Success 200 {array} R
// @Failure 500 {object} R
// @Router /products [get]
func (p *ProductController) GetAllProducts(ctx *gin.Context) {
	db := p.DB
	repository := repositories.NewProductRepo(db)
	products := repository.GetAllProducts()
	ctx.JSON(200, views.View(products))
}

// AddProduct godoc
// @Summary Add a new product
// @Description Adds a new product to the database
// @Tags Product
// @Accept  json
// @Produce  json
// @Param  product  body  models.Product  true  "Product JSON"
// @Success 200 {object} R
// @Failure 400 {object} R
// @Failure 500 {object} R
// @Router /products [post]
func (p *ProductController) AddProduct(ctx *gin.Context) {
	db := p.DB
	var product models.Product
	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, views.ErrView(status.ErrorCodeValidation, err.Error()))
		ctx.AbortWithStatus(500)
	}
	repository := repositories.NewProductRepo(db)
	inserted := repository.AddProduct(product)

	if inserted {
		ctx.JSON(200, views.View(nil))
		return
	} else {
		ctx.JSON(500, views.ErrView(status.ErrorCodeDB, "Product is not added"))
		return
	}
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Retrieves a product by its ID from the database
// @Tags Product
// @Produce  json
// @Param  id  path  int  true  "Product ID"
// @Success 200 {object} R
// @Failure 400 {object} R
// @Failure 404 {object} R
// @Failure 500 {object} R
// @Router /products/{id} [get]
func (p *ProductController) GetProductByID(ctx *gin.Context) {
	db := p.DB
	var ProductID models.ProductURI
	if err := ctx.ShouldBindUri(&ProductID); err != nil {
		ctx.JSON(http.StatusInternalServerError, views.ErrView(status.ErrorCodeValidation, err.Error()))
	}
	repository := repositories.NewProductRepo(db)
	product := repository.GetProductByID(ProductID.ID)
	if (product != models.Product{}) {
		ctx.JSON(200, views.View(product))
		return
	} else {
		ctx.JSON(http.StatusInternalServerError, views.ErrView(status.ErrorCodeDB, "Product is not obtained"))
	}
}

// DeleteProduct godoc
// @Summary Delete product by ID
// @Description Deletes a product from the database by its ID
// @Tags Product
// @Produce  json
// @Param  id  path  int  true  "Product ID"
// @Success 200 {object} R
// @Failure 400 {object} R
// @Failure 404 {object} R
// @Failure 500 {object} R
// @Router /products/{id} [delete]
func (p *ProductController) DeleteProduct(ctx *gin.Context) {
	db := p.DB
	var UserID models.ProductURI
	if err := ctx.ShouldBindUri(&UserID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	repository := repositories.NewProductRepo(db)
	deleted := repository.DeleteProduct(UserID.ID)
	if deleted {
		ctx.JSON(200, views.View(nil))
		return
	} else {
		ctx.JSON(300, views.ErrView(status.ErrorCodeDB, "Product is not deleted"))
		return
	}
}

// UpdateProduct godoc
// @Summary Update product by ID
// @Description Updates a product's details in the database by its ID
// @Tags Product
// @Accept json
// @Produce json
// @Param  id  path  int  true  "Product ID"
// @Param  product  body  models.Product  true  "Product JSON"
// @Success 200 {object} R
// @Failure 400 {object} R
// @Failure 404 {object} R
// @Failure 500 {object} R
// @Router /products/{id} [put]
func (p *ProductController) UpdateProduct(ctx *gin.Context) {
	db := p.DB
	var product models.Product
	var UserID models.ProductURI
	if err := ctx.ShouldBindUri(&UserID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	repository := repositories.NewProductRepo(db)
	product = repository.UpdateProduct(UserID.ID, product)
	if (product != models.Product{}) {
		ctx.JSON(200, gin.H{"status": "success", "product": product})
		return
	} else {
		ctx.JSON(http.StatusInternalServerError, views.ErrView(status.ErrorCodeDB, "Product is not updated"))
	}
}

// AddCategory godoc
// @Summary Add a category to a product
// @Description Associates a category with a product by their IDs
// @Tags Product
// @Accept json
// @Produce json
// @Param  product_id  path  int  true  "Product ID"
// @Param  category_id  path  int  true  "Category ID"
// @Success 200 {object} R
// @Failure 400 {object} R
// @Failure 404 {object} R
// @Failure 500 {object} R
// @Router /products/{product_id}/categories/{category_id} [post]
func (o *ProductController) AddCategory(c *gin.Context) {
	db := o.DB
	repository := repositories.NewCategoryRepository(db)
	var productID models.ProductURICategory
	var categoryID models.CategoryURI

	if err := c.ShouldBindUri(&productID); err != nil {
		c.JSON(http.StatusInternalServerError, views.ErrView(status.ErrorCodeValidation, "Product ID is not given"))
		return
	}

	if err := c.ShouldBindUri(&categoryID); err != nil {
		c.JSON(http.StatusInternalServerError, views.ErrView(status.ErrorCodeValidation, "Category ID is not given"))
		return
	}

	if added := repository.AddCategoryToProduct(productID.ID, categoryID.ID); added {
		c.JSON(http.StatusOK, views.View(nil))
	} else {
		c.JSON(http.StatusInternalServerError, views.ErrView(status.ErrorCodeDB, "Product is not added"))
	}
}

// DeleteCategory godoc
// @Summary Remove a category from a product
// @Description Disassociates a category from a product by their IDs
// @Tags Product
// @Accept json
// @Produce json
// @Param  product_id  path  int  true  "Product ID"
// @Param  category_id  path  int  true  "Category ID"
// @Success 200 {object} R
// @Failure 400 {object} R
// @Failure 404 {object} R
// @Failure 500 {object} R
// @Router /products/{product_id}/categories/{category_id} [delete]
func (o *ProductController) DeleteCategory(c *gin.Context) {
	db := o.DB
	repository := repositories.NewCategoryRepository(db)
	var productID models.ProductURICategory
	var categoryID models.CategoryURI

	if err := c.ShouldBindUri(&productID); err != nil {
		c.JSON(http.StatusInternalServerError, views.ErrView(status.ErrorCodeValidation, "Product ID is not given"))
		return
	}

	if err := c.ShouldBindUri(&categoryID); err != nil {
		c.JSON(http.StatusInternalServerError, views.ErrView(status.ErrorCodeValidation, "Category ID is not given"))
		return
	}

	if deleted := repository.DeleteCategoryProductRelationship(productID.ID, categoryID.ID); deleted {
		c.JSON(http.StatusOK, views.View(nil))
	} else {
		c.JSON(http.StatusInternalServerError, views.ErrView(status.ErrorCodeDB, "Category is not deleted"))
	}
}
