package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gods-mack/go-rest/controllers"
)

func UserRoute(router *gin.Engine) {

	router.GET("/", controllers.GetUsers)
	router.POST("/", controllers.CreateUser)

	router.GET("/user_by_id", controllers.GetUserById)

	router.GET("/dukaan", controllers.GetDukaan)
	router.POST("/dukaan", controllers.CreateDukaan)

	router.GET("/userprofile", controllers.GetUserProfile)
	router.POST("/userprofile", controllers.CreateUserProfile)



}
