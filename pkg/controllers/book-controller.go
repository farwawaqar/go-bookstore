package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/farwawaqar/go-book-store/pkg/models/book"
	"github.com/farwawaqar/go-book-store/pkg/utils"
	"github.com/farwawaqar/go-bookstore/pkg/models"
	"github.com/farwawaqar/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book
func GetBooks(w http.ResponseWriter, r *http.Request)
{
	newBooks:=models.GetAllBooks()
	res, _ :=json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r) //we will get access to r -Request. inside request we have bookid
	bookId :=vars["bookId"]
	ID,err := strconv.ParseInt(bookId,0,0)
	if err !=nil{
		fmt.Println("error while parsing")
	}
	bookDetails, _ :=models.GetBookById(ID)
	res, _:= json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter,r *http.Request){
	CreateBook:=&models.Book{}
	utils.ParseBody(r,CreateBook)
	b:=CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Delete(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId:= vars["bookId"]
	//we need to access the reqeest
	//we need to access the  bookid
	//then we need to convert into int
	//then if we have err we will print it out
	ID,err := strconv.ParseInt(bookId,0,0)
	if err !=nil{
		fmt.Println("error while parsing")
	}
	book :=models.DeleteBook(ID)
	res,_ :=json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request)
{
	var updateBook =&models.Book{}
	utils.ParseBody(r,updateBook) // whatever is in the request -id 
	//we will parse it and unmarshall it into the format that golang or database understands
	vars:= mux.Vars(r)
	bookId:= vars["bookId"]
	ID,err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails,db:=models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name =updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author=updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication =updateBook.Publication
	}
	db.Save(&bookDetails)
	res,_ :=json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}