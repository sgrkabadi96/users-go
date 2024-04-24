package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sgrkabadi96/bookAppClone/config"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name string `json:"name"`
	Age int `json:"age"`
	Website string `json:"Website"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}


func(user *User) CreateUser() {
	db.Create(user)
}

func GetAllUsers() []User {
	var users []User
	db.Find(&users)
	return users
}

func GetUserById(id int) User {
	var user User 
	db.Where("ID=?" , id ).Find(&user)
	return user 
}

func DeleteById(id int) User {
	var user User 
	db.Where("ID=?" , id).Delete(&user)
	return user
}

func DeleteAllUser() []User {
	var users []User
	db.Delete(&users)
	return users
}


