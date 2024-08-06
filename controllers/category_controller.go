package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/muhammadjon1304/e-commerce/models"
	"github.com/muhammadjon1304/e-commerce/repositories"
	"net/http"
)

type CategoryController struct {
	DB *sql.DB
}

func NewCategoryController(db *sql.DB) CategoryController {
	return CategoryController{DB: db}
}

func (c *CategoryController) GetAllCategory(ctx *gin.Context) {
	db := c.DB
	repository := repositories.NewCategoryRepository(db)
	categories := repository.GetAllCategories()
	ctx.JSON(200, gin.H{"categories": categories})
}

func (c *CategoryController) AddCategory(ctx *gin.Context) {
	db := c.DB
	var Category models.Category
	if err := ctx.ShouldBindJSON(&Category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	repository := repositories.NewCategoryRepository(db)
	inserted := repository.AddCategory(Category)
	if inserted {
		ctx.JSON(200, gin.H{"status": "success", "msg": "category created"})
	} else {
		ctx.JSON(500, gin.H{"status": "fail", "msg": "category not created"})
		return
	}
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	db := c.DB
	var CategoryID models.CategoryURI
	if err := ctx.ShouldBindUri(&CategoryID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	repository := repositories.NewCategoryRepository(db)
	deleted := repository.DeleteCategory(CategoryID.ID)
	if deleted {
		ctx.JSON(200, gin.H{"status": "success", "message": "category deleted"})
		return
	} else {
		ctx.JSON(300, gin.H{"status": "fail", "message": "can't delete category"})
		return
	}
}
