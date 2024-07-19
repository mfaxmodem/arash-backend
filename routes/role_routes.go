package routes

import (
	"arash-website/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoleRoutes(router *gin.RouterGroup, roleController *controllers.RoleController) {
	roleRoutes := router.Group("/roles")
	{
		roleRoutes.POST("/", roleController.CreateRole)
		roleRoutes.GET("/", roleController.GetAllRoles)
		roleRoutes.GET("/:id", roleController.GetRoleByID)
		roleRoutes.PUT("/:id", roleController.UpdateRole)
		roleRoutes.DELETE("/:id", roleController.DeleteRole)
	}
}
