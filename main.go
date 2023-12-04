package main

import (
	"github.com/iqbaltc13/back-end-learner-tryout-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error":   false,
			"message": "Yayyyy I'am Gin Gonic",
		})
	})

	public := r.Group("/api")

	public.POST("/auth/register", controllers.Register)

	r.Run()
}
