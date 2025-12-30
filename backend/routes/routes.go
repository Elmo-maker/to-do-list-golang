package routes

import (
	"github.com/gin-gonic/gin"
	"to-do-list/controllers"
	"to-do-list/middleware"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/todos", controllers.GetTodos)
		api.GET("/todos/:id", controllers.GetTodoByID)
		api.POST("/todos", controllers.CreateTodo)
		api.PUT("/todos/:id", controllers.UpdateTodo)
		api.DELETE("/todos/:id", controllers.DeleteTodo)
	}
}
