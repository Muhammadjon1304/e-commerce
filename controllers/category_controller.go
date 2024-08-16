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

type CategoryController struct {
	DB *sql.DB
}

func NewCategoryController(db *sql.DB) CategoryController {
	return CategoryController{DB: db}
}

// GetAllCategory godoc
// @Summary Get all categories
// @Description Retrieves all categories from the database
// @Tags Category
// @Accept  json
// @Produce  json
// @Success 200 {object} R
// @Failure 500 {object} R
// @Router /categories [get]
func (c *CategoryController) GetAllCategory(ctx *gin.Context) {
	db := c.DB
	repository := repositories.NewCategoryRepository(db)
	categories := repository.GetAllCategories()
	ctx.JSON(200, views.View(categories))
}

// AddCategory godoc
// @Summary Add a new category
// @Description Creates a new category in the database
// @Tags Category
// @Accept  json
// @Produce  json
// @Param   category  body  models.Category  true  "Category JSON"
// @Success 200 {object} R
// @Failure 400 {object} R
// @Failure 500 {object} R
// @Router /categories [post]
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
		ctx.JSON(200, views.ErrView(status.ErrorCodeDB, "Category not added"))
	} else {
		ctx.JSON(500, views.View(nil))
		return
	}
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Deletes a category from the database by category ID
// @Tags Category
// @Accept  json
// @Produce  json
// @Param   categoryID   path  int  true  "Category ID"
// @Success 200 {object} map[string]string  R
// @Failure 400 {object} map[string]string  R
// @Failure 500 {object} map[string]string  R
// @Router /categories/{categoryID} [delete]
func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	db := c.DB
	var CategoryID models.CategoryURI
	if err := ctx.ShouldBindUri(&CategoryID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	repository := repositories.NewCategoryRepository(db)
	deleted := repository.DeleteCategory(CategoryID.ID)
	if deleted {
		ctx.JSON(200, views.View(nil))
		return
	} else {
		ctx.JSON(300, views.ErrView(status.ErrorCodeDB, "Category is not deleted"))
		return
	}
}
