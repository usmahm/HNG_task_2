package routes

import (
	"person-app/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/api", controllers.CreateUser)
	router.GET("/api/:user_id", controllers.GetUser)
	router.PATCH("/api/:user_id", controllers.UpdateUser)
	router.DELETE("/api/:user_id", controllers.DeleteUser)
}
