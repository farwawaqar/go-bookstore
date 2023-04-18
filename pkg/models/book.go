package models

import (
	"github.com/farwawaqar/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	//Id          string `json:"id"`
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {

	//we need to intialize the database
	config.Connect()
	db = config.GetDB() //getting the db from config file
	db.AutoMigrate(&Book{})
}

/*In Go, a function receiver is a way of attaching a function to a particular type or struct. It allows you to define methods on types or structs, which can then be called on instances of those types or structs.

A function receiver is defined as the first parameter in a method signature, and it specifies the type or struct that the method is attached to. It can be either a value receiver or a pointer receiver.

A value receiver is defined using the type itself, while a pointer receiver is defined using a pointer to the type. A value receiver works with a copy of the instance, while a pointer receiver works with a reference to the instance, allowing you to modify its values.

For example, in the code line func (b *Book) CreateBook() *Book, (b *Book) is the function receiver. It specifies that the CreateBook() method is associated with the Book type and takes a pointer to a Book instance as its receiver. This means that the CreateBook() method can modify the state of the Book instance it is called on.

In general, function receivers are a powerful feature of Go that allows you to define methods on types or structs and create more expressive and reusable code.

*/

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

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

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(&book)
	return book
}
