package main

import (
	"fmt"
	"log"
	"os"
	"person-app/database"
	"person-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	fmt.Println("dduri", uri)

	database.MongoClient = database.DBInstance()

	router := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	routes.UserRoutes(router)

	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
