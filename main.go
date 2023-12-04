package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/iqbaltc13/back-end-learner-tryout-api/controllers"
	"github.com/iqbaltc13/back-end-learner-tryout-api/database"
	"github.com/iqbaltc13/back-end-learner-tryout-api/middleware"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()

}
func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	//database.Database.AutoMigrate(&model.User{})
	//database.Database.AutoMigrate(&model.Entry{})
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/auth/register", controllers.Register)
	publicRoutes.POST("/auth/login", controllers.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controllers.AddEntry)
	protectedRoutes.GET("/entry", controllers.GetAllEntries)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
