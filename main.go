package main

import (
	"log"
	"os"
	"person-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
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
