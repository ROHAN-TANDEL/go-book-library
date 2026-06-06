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

func getAuthor(c *gin.Context) {
	var authorId = c.Param("id")
	author, err := getAuthorRepo(authorId)

	if err != nil {
		c.JSON(400, gin.H{"data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": author})
}

func getAuthors(c *gin.Context) {
	authors, err := getAuthorsRepo()

	if err != nil {
		c.JSON(400, gin.H{"data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": authors})
}

func addAuthor(c *gin.Context) {
	var inputAuthor Author
	if err := c.ShouldBindJSON(&inputAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := addAuthorRepo(inputAuthor)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func replaceAuthor(c *gin.Context) {
	var authorId = c.Param("id")
	var inputAuthor Author
	if err := c.ShouldBindJSON(&inputAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := replaceAuthorRepo(authorId, inputAuthor)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func upgradeAuthor(c *gin.Context) {

	var inputAuthor AuthorPatch

	if err := c.ShouldBindJSON(&inputAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := upgradeAuthorRepo(c.Param("id"), inputAuthor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func removeAuthor(c *gin.Context) {
	var authorId = c.Param("id")
	res, err := removeAuthorRepo(authorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Author does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": authorId})
}

func getCategory(c *gin.Context) {
	var catID = c.Param("id")

	res, err := getCategoryRepo(catID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
func getCategories(c *gin.Context) {
	res, err := getCategoriesRepo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}
func addCategory(c *gin.Context) {
	var category Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := addCategoryRepo(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
func replaceCategory(c *gin.Context) {
	var category Category
	var catID = c.Param("id")
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := replaceCategoryRepo(catID, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}
func upgradeCategory(c *gin.Context) {
	var category CategoryPatch
	var catID = c.Param("id")
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := upgradeCategoryRepo(catID, category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
func removeCategory(c *gin.Context) {
	var catID = c.Param("id")
	res, err := removeCategoryRepo(catID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
