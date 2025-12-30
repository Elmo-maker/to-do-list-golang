package main

import (
	"github.com/gin-gonic/gin"
	"to-do-list/config"
	"to-do-list/routes"
)

func main () {
	config.ConnectDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "SERVER RUNNING KINGS!",
		})
	})

	routes.SetupRoutes(r)
	
	r.Run(":8080")
}