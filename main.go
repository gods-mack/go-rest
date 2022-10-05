package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gods-mack/go-rest/routes"

	"github.com/gods-mack/go-rest/config"
)

func testing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// main

// run project
// go run main.go
// docker-compose up --build   (to run postgres)
func main() {
	router := gin.New()

	config.Connect()

	routes.UserRoute(router)

	router.GET("/testing", testing)

	router.Run(":8005")
}
