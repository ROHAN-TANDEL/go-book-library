package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addBook(c *gin.Context) {
	var record newBook
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := addBookRepo(&record)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func getBook(c *gin.Context) {

	var bookId = c.Param("id")

	books, err := getBookRepo(bookId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func getBooks(c *gin.Context) {

	books, err := getBooksRepo()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func removeBook(c *gin.Context) {
	var bookId = c.Param("id")

	rowsAffected, err := removeBookRepo(bookId)

	if err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Book does not exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookId})
}

func replaceBook(c *gin.Context) {
	var record newBook
	var bookID = c.Param("id")
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record, err := replaceBookRepo(bookID, record)

	if err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": record})
}

func upgradeBook(c *gin.Context) {
	var record newBook
	var bookId = c.Param("id")
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := upgradeBookRepo(bookId, record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Book does not exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": record})
}
