package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/books", ApiGetBooks)
	app.GET("/books/:id", GetBookById)
	app.POST("/books", AddBooks)
	app.Run("localhost:1337")
}
