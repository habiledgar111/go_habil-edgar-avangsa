package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id   int    `json:"ID" from:"id"`
	Name string `json:"name" from:"name"`
	Age  int    `json:"age" from:"age"`
}

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "users"
}
