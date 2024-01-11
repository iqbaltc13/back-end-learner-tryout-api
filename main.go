package main

import (
	"fmt"
	"log"

	"os"

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
	err := godotenv.Load(".env")
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
	protectedRoutes := router.Group("/api")
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("register", controllers.Register)
	publicRoutes.POST("login", controllers.Login)

	protectedRoutes.Use(middleware.JWTAuthMiddleware())

	port := ":" + os.Getenv("PORT")
	router.Run(port)
	fmt.Println("Server running on port " + port)
}
