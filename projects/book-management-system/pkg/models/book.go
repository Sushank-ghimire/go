package models

import (
	"book-management-system/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

type Book struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`
	Author      string `gorm:"column:author" json:"author"`
	Publication string `gorm:"column:publication" json:"publication"`
}

func init() {
	config.ConnectDatabase()
	db := config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) createBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var bookNeeded Book
	db_book := db.Where("ID=?", id).Find(&bookNeeded)
	return &bookNeeded, db_book
}

func DeleteBook(id int64) Book {
	var DeletedBook Book
	db.Where("ID=?", id).Delete(&DeletedBook)
	return DeletedBook
}
