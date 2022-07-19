package server

import (
	"bookstore/handler"

	"github.com/gin-gonic/gin"
)

// runs the book server on port 8080 of localhost
func Run() error {
	router := gin.Default()
	router.GET("/book", handler.GetBook)
	router.GET("/book/:id", handler.GetBookByID)
	router.POST("/book", handler.CreateBook)

	router.Run("localhost:8080")

	return nil
}
