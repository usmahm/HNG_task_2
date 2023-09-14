package routes

import (
	"person-app/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/api", controllers.CreatePerson)
	router.GET("/api/:param", controllers.GetPerson)
	router.PUT("/api/:param", controllers.UpdatePerson)
	router.DELETE("/api/:param", controllers.DeletePerson)
}
