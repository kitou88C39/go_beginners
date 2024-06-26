package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct{
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Quantity int  `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In search of Lost Time", Author: "Marcel Proust", Quantity: 2 },
	{ID: "2", Title: "The Great Gatsby", Author: "Scott Fitzgerald", Quantity: 5 },
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6 },
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context){
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"messeage":"Book not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func checkoutBooks(c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"messeage":"Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"messeage":"Book not found."})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"messeage":"Book not available."})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context){
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"messeage":"Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"messeage":"Book not found."})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string)(*book, error){
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func creatBooks(c *gin.Context){
	var newBook book
	if err := c.BindJSON(&newBook); err != nil{
		return
	}
	books = append(books,newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}


func main(){
	router := gin.Default()
	router.GET("/books",getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", creatBooks)
	router.PATCH("/checkout", checkoutBooks)
	router.PATCH("/retrurn", returnBook)
	router.Run("localhost:8080")
} 