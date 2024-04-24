package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
)

var (
	db *gorm.DB
)


func Connect() {
	// Print the constructed connection string
	
	dsn := fmt.Sprintf("{username}:{password}@tcp(localhost:3306)/userApp?charset=utf8&parseTime=True")

	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err) // Print the error message
	}
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db;
}