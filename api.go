package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiGetBooks(c *gin.Context) {

	c.JSON(http.StatusOK, listofBooks)
}

func AddBooks(c *gin.Context) {
	var data books

	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, apiresponse{Message: "Body Is Missing!"})
		return
	}
	if data.Price <= 0 {
		c.JSON(http.StatusBadRequest, apiresponse{Message: "Price is not set!"})
		return
	}
	if data.Title == "" {
		c.JSON(http.StatusBadRequest, apiresponse{Message: "Title is not set!"})
		return
	}
	listofBooks = append(listofBooks, data)
	c.JSON(http.StatusCreated, data)
}

func GetBookById(c *gin.Context) {
	index := c.Param("id")
	for _, book := range listofBooks {
		if book.ID == index {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	c.JSON(http.StatusNotFound, apiresponse{Message: "Book not found!"})

}
