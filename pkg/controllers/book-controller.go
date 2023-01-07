package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ashwin/go-mysql/pkg/models"
	"github.com/gin-gonic/gin"
)

var NewBook models.Book

func GetBook(c *gin.Context) {
	newBooks := models.GetAllBooks()
	c.IndentedJSON(http.StatusOK, newBooks)
}

func GetBookById(c *gin.Context) {
	Id := c.Param("id")
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	book, _ := models.GetBookById(ID)
	c.IndentedJSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {

	var newBook models.Book
	err := c.BindJSON(&newBook)
	if err != nil {
		log.Fatal(err)
	}
	b := newBook.CreateBook()
	c.IndentedJSON(http.StatusOK, b)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	book := models.DeleteBook(ID)
	c.IndentedJSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	var updateBook models.Book
	err := c.BindJSON(&updateBook)
	if err != nil {
		log.Fatal(err)
	}
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)

	c.IndentedJSON(http.StatusOK, bookDetails)
}
