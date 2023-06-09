package model

import "gorm.io/gorm"

type Kategori struct {
	gorm.Model
	ID     int    `gorm:"primary_key;not null"`
	Nama   string `json:"kategori" form:"kategori" gorm:"type:varchar(255)"`
	Barang []Barang
}
