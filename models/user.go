package models

import "gorm.io/gorm"

// import (
// 	"fmt"
// 	"log"
// 	"time"
// )

type User struct {
	gorm.Model
	Id      			int    `json:"id" gorm:"primary_key"`
	UserHandle 			string `json:"user_handle" gorm:"uniqueIndex"`
	Name     			string `json:"name"`
	Email    			string `json:"email"`
	Password 			string `json:"password"`
	FollowerCount 		int 	`json:"follower_count"`
	FollowingCount 		int `json:"following_count"`
	ProfileDP 			string `json:"profile_dp"`
	PostCount 			int 	`json:"post_count"`
	Bio 				string `json:"bio"`
}

type Dukaan struct {
	gorm.Model
	Id      int    `json:"dukaan_id" gorm:"primary_key"`
	Name    string `json:"dukaan_name"`
	OwnerId int    `json:"owner_id"`
	Owner   User   `gorm:"foreignKey:OwnerId"`
	// Foreign Key
}

type UserProfile struct {
	gorm.Model
	profile_id int `json:"profile_id" gorm:"primary_key`	
	user User 	`gorm:"foreignKey:Id"`
}

type UserRelationship struct {
	gorm.Model
	Id int `json:"id" gorm:"primary_key"`
	Followers []User `gorm:"many2many:user_followers;"`
	Followings []User `gorm:"many2many:user_followings;"`
	user User `gorm:"foreignKey:Id"`
	
}