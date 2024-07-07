package controllers

import (
	"arash-website/models"
	"arash-website/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryController struct {
	CategoryService *services.CategoryService
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var category models.Category
	if err := ctx.BindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.CategoryService.CreateCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	ctx.JSON(http.StatusCreated, category)
}

// دیگر متدهای مربوط به خواندن، به روزرسانی و حذف دسته‌بندی‌ها را اضافه کنید
