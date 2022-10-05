package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gods-mack/go-rest/config"
	"github.com/gods-mack/go-rest/models"
	//"fmt"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, gin.H{
		"result": &users,
		"count":  len(users),
	})

}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	x := config.DB.Create(&user)
	
	if x.Error == nil {
		c.JSON(200, gin.H{
			"result": &user,
			"msg": "success",
		})
		return
	}
	c.JSON(400, gin.H{
			"msg":"user handle not valid",
	})
}

func GetUserById(c *gin.Context) {
	user_id := c.Query("user_id")
	user := models.User{}
	config.DB.Find(&user, user_id)

	c.JSON(200, gin.H{
		"result": &user,
		"msg":"success",
	})	
	
}


func DeleteUser(c *gin.Context) {
	//user_id := c.Params.ByName("id")
	//config.DB.where("id = ?", user_id).Delete(&models.User{})
	//c.JSON(200, gin.H{"id #" + user_id: "deleted"})
}

func GetDukaan(c *gin.Context) {
	dukaans := []models.Dukaan{}
	config.DB.Find(&dukaans)
	c.JSON(200, gin.H{
		"result": &dukaans,
		"count":  len(dukaans),
	})

}

func CreateDukaan(c *gin.Context) {

	var dukaan models.Dukaan
	c.BindJSON(&dukaan)
	config.DB.Create(&dukaan)
	c.JSON(200, &dukaan)
}




func GetUserProfile(c *gin.Context) {

	profiles :=  []models.UserProfile{}
	config.DB.Find(&profiles)
	c.JSON(200, gin.H{
		"result": &profiles,
		"count":  len(profiles),
	})
}



func CreateUserProfile(c *gin.Context) {

	var profile models.UserProfile
	c.BindJSON(&profile)
	config.DB.Create(&profile)
	c.JSON(200, &profile)
}
