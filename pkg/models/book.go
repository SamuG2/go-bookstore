package models

import (
	"fmt"

	"github.com/SamuG2/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func CreateBook(b *Book) *Book {
	//db.NewRecord(b)
	db.Create(&b)
	return b
}

// func (b *Book) CreateBook() *Book {
// 	db.NewRecord(b)
// 	db.Create(&b)
// 	return b
// }

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	fmt.Println("ID:", Id)
	db.Where("ID=?", Id).Delete(&book)
	fmt.Printf("%+v", book)
	return book
}
