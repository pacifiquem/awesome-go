package main


import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct{
	ID 			string 	`json:"id"`
	Title		string 	`json:"title"`
	Author 		string 	`json:"author"`
	Quantity 	int 	`json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Who we are?", Author: "M.Pac", Quantity: 3},
	{ID: "2", Title: "Peace and Love", Author: "Murangwa", Quantity: 5},
	{ID: "1", Title: "The lost memory", Author: "Pacifique", Quantity: 0},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func returnBookById(id string) (*book, error) {
	for i, book := range books {
		if book.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("Book not found")
}

func getBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := returnBookById(id)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)	

}

func main() {
	router := gin.Default();
	router.GET("/books", getBooks)
	router.POST("/new/book", createBook)
	router.GET("/book/:id", getBookById)
	router.Run("localhost:8080")
}