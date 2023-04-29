package config

import (
	"sec024/praktikum/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connection() {
	dsn := "root:Mbahbambang123@tcp(localhost:3306)/sec21orm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}

func Migration() error {
	return DB.AutoMigrate(entity.User{})
}
