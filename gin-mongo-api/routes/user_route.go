package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	//All routes related to users comes here
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetAUser())
	router.PUT("/user/:userId/pinreset", controllers.EditAUser())
	router.PUT("/user/:userId/deposit", controllers.DepositAmount())
	router.GET("/users", controllers.GetAllUsers())
}
