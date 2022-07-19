package handler

import (
	"bookstore/model"
	"bookstore/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBooks responds with a list of books in the repository
func GetBook(c *gin.Context) {
	if r, err := repository.NewRepository(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			gin.H{"message": "invalid repository"})
	} else {
		c.IndentedJSON(http.StatusOK, r.GetBooks())
	}
}

// GetBookByID responds with a specific book from the library
func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	if r, err := repository.NewRepository(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			gin.H{"message": "invalid repository"})
	} else {
		if b, err := r.FindBookByID(id); err != nil {
			c.IndentedJSON(http.StatusOK, b)
		} else {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		}
	}
}

// CreateBook adds a book to the repository of books
func CreateBook(c *gin.Context) {
	var newBook model.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	if r, err := repository.NewRepository(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			gin.H{"message": "invalid repository"})
	} else {
		r.AddBook(newBook)
		c.IndentedJSON(http.StatusCreated, newBook)
	}
}
