package routes

import (
	"arash-website/controllers"
	"github.com/gin-gonic/gin"
)

func SetupCommentRoutes(router *gin.RouterGroup, commentController *controllers.CommentController) {
	commentRoutes := router.Group("/comments")
	{
		commentRoutes.POST("/", commentController.CreateComment)
		commentRoutes.GET("/", commentController.GetAllComments)
		commentRoutes.GET("/:id", commentController.GetCommentByID)
		commentRoutes.PUT("/:id", commentController.UpdateComment)
		commentRoutes.DELETE("/:id", commentController.DeleteComment)
	}
}
