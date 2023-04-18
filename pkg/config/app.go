package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //"_ is nesseccary"
)

//create variable db that will help other file to connect with db

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "username/password/tablename??charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
