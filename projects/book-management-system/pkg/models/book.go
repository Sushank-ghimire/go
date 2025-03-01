package models

import (
	"book-management-system/pkg/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`
	Author      string `gorm:"column:author" json:"author"`
	Publication string `gorm:"column:publication" json:"publication"`
}

func init() {
	config.ConnectDatabase()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// ✅ Fix: Ensure the database connection is used correctly
func (b *Book) CreateBook() *Book {
	if db == nil {
		fmt.Println("Database connection is nil")
		return nil
	}
	db.Create(&b)
	return b
}

// ✅ Fix: Handle database errors properly
func GetAllBooks() []Book {
	var Books []Book
	if db == nil {
		fmt.Println("Database connection is nil")
		return []Book{}
	}
	result := db.Find(&Books)
	if result.Error != nil {
		fmt.Println("Error fetching books:", result.Error)
		return []Book{}
	}
	return Books
}

func GetBookById(id int64) (*Book, error) {
	var bookNeeded Book
	result := db.Where("ID = ?", id).First(&bookNeeded)

	if result.Error != nil {
		if result.RecordNotFound() {
			fmt.Println("Book not found in database") // Debug log
			return nil, fmt.Errorf("book not found")
		}
		fmt.Println("Database error:", result.Error) // Debug log
		return nil, result.Error
	}

	fmt.Println("Book found:", bookNeeded) // Debug log
	return &bookNeeded, nil
}

// ✅ Fix: Ensure deleted book is correctly returned
func DeleteBook(id int64) Book {
	var deletedBook Book
	db.Where("ID = ?", id).First(&deletedBook)
	db.Delete(&deletedBook)
	return deletedBook
}
