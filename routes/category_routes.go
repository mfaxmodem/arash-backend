package routes

import (
	"arash-website/models"
	"arash-website/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SetupCategoryRoutes(router *gin.Engine, categoryService *services.CategoryService) {
	router.POST("/categories", createCategory(categoryService))
	router.GET("/categories/:id", getCategory(categoryService))
	router.PUT("/categories/:id", updateCategory(categoryService))
	router.DELETE("/categories/:id", deleteCategory(categoryService))
	router.GET("/categories", getAllCategories(categoryService))
}

func createCategory(categoryService *services.CategoryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var category models.Category
		if err := c.ShouldBindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := categoryService.CreateCategory(&category); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, category)
	}
}

func getCategory(categoryService *services.CategoryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
			return
		}
		category, err := categoryService.GetCategoryByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		c.JSON(http.StatusOK, category)
	}
}

func updateCategory(categoryService *services.CategoryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
			return
		}
		var category models.Category
		if err := c.ShouldBindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category data"})
			return
		}
		category.ID = uint(id)
		if err := categoryService.UpdateCategory(&category); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, category)
	}
}

func deleteCategory(categoryService *services.CategoryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
			return
		}
		var category models.Category
		category.ID = uint(id)
		if err := categoryService.DeleteCategory(&category); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{})
	}
}

func getAllCategories(categoryService *services.CategoryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := categoryService.GetAllCategories()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, categories)
	}
}
