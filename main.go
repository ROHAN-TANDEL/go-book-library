package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/get-book", getBook)
	router.POST("/add-book", addBook)
	router.PUT("/replace-book", replaceBook)
	router.PATCH("/upgrade-book", upgradeBook)
	router.DELETE("/remove-book", removeBook)
	router.Run(":8080")
}

func getBook(c *gin.Context) {
	c.JSON(200, gin.H{"action": "book is fetched"})
}

func addBook(c *gin.Context) {
	c.JSON(200, gin.H{"action": "book is added"})
}

func replaceBook(c *gin.Context) {
	c.JSON(200, gin.H{"action": "book is replaced"})
}

func upgradeBook(c *gin.Context) {
	c.JSON(200, gin.H{"action": "book is upgraded"})
}

func removeBook(c *gin.Context) {
	c.JSON(200, gin.H{"action": "book is removed"})
}
