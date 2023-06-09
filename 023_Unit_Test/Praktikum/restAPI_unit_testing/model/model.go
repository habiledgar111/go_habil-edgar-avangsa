package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `gorm:"primary_key;not null"`
	Email    string `gorm:";type:varchar(255)unique;not null"`
	Password string `gorm:"notnull"`
	Name     string `gorm:"type:varchar(255)"`
	Age      int    `gorm:"type:int"`
}
type UserMock struct {
	gorm.Model
	Email    string
	Password string
	Name     string
	Age      int
}

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "users"
}
