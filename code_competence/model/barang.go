package model

import (
	"gorm.io/gorm"
)

type Barang struct {
	gorm.Model
	ID        int     `gorm:"primary_key;not null"`
	Nama      string  `json:"nama" form:"nama" gorm:"type:varchar(255)"`
	Deskripsi string  `json:"deskripsi" form:"deskripsi" gorm:"type:varchar(255)"`
	Jumlah    int     `json:"jumlah" form:"harga" gorm:"type:int"`
	Harga     float64 `json:"harga" form:"harga" gorm:"type:double"`
	Kategori  Kategori
}
