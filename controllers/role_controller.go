package controllers

import (
	"arash-website/config"
	"arash-website/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&role)
	c.JSON(http.StatusOK, role)
}

func GetAllRoles(c *gin.Context) {
	var roles []models.Role
	config.DB.Find(&roles)
	c.JSON(http.StatusOK, roles)
}

func GetRoleByID(c *gin.Context) {
	var role models.Role
	id := c.Param("id")
	if err := config.DB.Where("id = ?", id).First(&role).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusOK, role)
}

func UpdateRole(c *gin.Context) {
	var role models.Role
	id := c.Param("id")
	if err := config.DB.Where("id = ?", id).First(&role).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&role)
	c.JSON(http.StatusOK, role)
}

func DeleteRole(c *gin.Context) {
	var role models.Role
	id := c.Param("id")
	if err := config.DB.Where("id = ?", id).Delete(&role).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role deleted"})
}
