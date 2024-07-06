package routes

import (
	"arash-website/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoleRoutes(router *gin.Engine) {
	roleRoutes := router.Group("/roles")
	{
		roleRoutes.POST("/", controllers.CreateRole)
		roleRoutes.GET("/", controllers.GetAllRoles)
		roleRoutes.GET("/:id", controllers.GetRoleByID)
		roleRoutes.PUT("/:id", controllers.UpdateRole)
		roleRoutes.DELETE("/:id", controllers.DeleteRole)
	}
}
