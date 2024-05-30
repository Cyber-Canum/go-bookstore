package models

import (
	"github.com/Cyber-Canum/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var b []Book
	db.Find(&b)
	return b
}

func GetBookByID(id int64) (*Book, *gorm.DB) {
	var b Book
	db := db.Where("ID=?", id).Find(&b)
	return &b, db
}

func DeleteBook(id int64) Book {
	var b Book
	db.Where("ID=?", id).Delete(b)
	return b
}
