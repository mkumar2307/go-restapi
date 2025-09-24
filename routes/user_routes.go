package routes

import (
	"github.com/mkumar2307/go-restapi/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)         // List all users
		userRoutes.POST("/", controllers.CreateUser)      // Create new user
		userRoutes.GET("/:id", controllers.GetUser)       // Get single user
		userRoutes.PUT("/:id", controllers.UpdateUser)    // Update user
		userRoutes.DELETE("/:id", controllers.DeleteUser) // Delete user
	}
}
