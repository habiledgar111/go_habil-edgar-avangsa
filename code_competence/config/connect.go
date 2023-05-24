package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Open() error {
	dsn := "root:Mbahbambang123@tcp(localhost:3306)/codecompetence?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	if err != nil {
		return err
	}
	return nil
}
