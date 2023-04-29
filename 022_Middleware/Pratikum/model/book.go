package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id       int    `json:"ID" from:"id"`
	Judul    string `json:"judul" from:"judul"`
	Penulis  string `json:"penulis" from:"penulis"`
	Penerbit string `json:"penerbit" from:"penerbit"`
}

func (Book) TableName() string {
	return "books"
}
