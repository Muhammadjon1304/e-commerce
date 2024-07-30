package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/muhammadjon1304/e-commerce/models"
	"github.com/muhammadjon1304/e-commerce/repositories"
	"net/http"
)

type ProductController struct {
	DB *sql.DB
}

func NewProductController(db *sql.DB) ProductController {
	return ProductController{DB: db}
}

func (p *ProductController) GetAllProducts(ctx *gin.Context) {
	db := p.DB
	repository := repositories.NewProductRepo(db)
	products := repository.GetAllProducts()
	ctx.JSON(200, gin.H{"products": products})
}

func (p *ProductController) AddProduct(ctx *gin.Context) {
	db := p.DB
	var product models.Product
	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		ctx.AbortWithStatus(500)
	}
	repository := repositories.NewProductRepo(db)
	inserted := repository.AddProduct(product)

	if inserted {
		ctx.JSON(200, gin.H{"status": "success", "msg": "product created"})
		return
	} else {
		ctx.JSON(500, gin.H{"status": "fail", "msg": "product not created"})
		return
	}
}

func (p *ProductController) GetProductByID(ctx *gin.Context) {
	db := p.DB
	var ProductID models.ProductURI
	if err := ctx.ShouldBindUri(&ProductID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	repository := repositories.NewProductRepo(db)
	product := repository.GetProductByID(ProductID.ID)
	if (product != models.Product{}) {
		ctx.JSON(200, gin.H{"status": "success", "product": product})
		return
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "error": "Can't get product"})
	}
}

func (p *ProductController) DeleteProduct(ctx *gin.Context) {
	db := p.DB
	var UserID models.ProductURI
	if err := ctx.ShouldBindUri(&UserID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	repository := repositories.NewProductRepo(db)
	deleted := repository.DeleteProduct(UserID.ID)
	if deleted {
		ctx.JSON(200, gin.H{"status": "success", "message": "product deleted"})
		return
	} else {
		ctx.JSON(300, gin.H{"status": "fail", "message": "can't delete product"})
		return
	}
}

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
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "error": "Can't get product"})
	}
}
