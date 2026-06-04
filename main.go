package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/get-book/:id", getBook)
	router.POST("/add-book/:id", addBook)
	router.PUT("/replace-book/:id", replaceBook)
	router.PATCH("/upgrade-book/:id", upgradeBook)
	router.DELETE("/remove-book/:id", removeBook)
	router.Run(":8080")
}

func getBook(c *gin.Context) {
	bookID := c.Param("id")
	c.JSON(200, gin.H{"action": "book is fetched", "book_id": bookID})
}

func addBook(c *gin.Context) {
	bookID := c.Param("id")
	c.JSON(200, gin.H{"action": "book is added", "book_id": bookID})
}

func replaceBook(c *gin.Context) {
	bookID := c.Param("id")
	c.JSON(200, gin.H{"action": "book is replaced", "book_id": bookID})
}

func upgradeBook(c *gin.Context) {
	bookID := c.Param("id")
	c.JSON(200, gin.H{"action": "book is upgraded", "book_id": bookID})
}

func removeBook(c *gin.Context) {
	bookID := c.Param("id")
	c.JSON(200, gin.H{"action": "book is removed", "book_id": bookID})
}
