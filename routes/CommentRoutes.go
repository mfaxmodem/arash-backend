package routes

import (
	"arash-website/controllers"
	"github.com/gin-gonic/gin"
)

func SetupCommentRoutes(r *gin.Engine, commentController *controllers.CommentController) {
	comments := r.Group("/comments")
	{
		comments.POST("/", commentController.CreateComment)
		comments.GET("/product/:productID", commentController.GetCommentsByProductID)
		comments.GET("/", commentController.GetAllComments)
		comments.PUT("/approve/:id", commentController.ApproveComment)
	}
}
