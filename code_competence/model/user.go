package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `gorm:"primary_key;not null"`
	Email    string `json:"email" form:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" form:"password" gorm:"type:varchar(255)"`
}
