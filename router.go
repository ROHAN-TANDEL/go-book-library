package main

import "github.com/gin-gonic/gin"

func route(router *gin.Engine) *gin.Engine {

	router = gin.Default()

	//book routers
	book := router.Group("/book")
	{
		book.GET("/:id", getBook)
		book.GET("/allwertyuiolkjhgvfcxdfghjmnbvcfghjk", getBooks)
		book.POST("/add", addBook)
		book.PUT("/:id", replaceBook)
		book.PATCH("/:id", upgradeBook)
		book.DELETE("/:id", removeBook)
	}

	author := router.Group("/author")
	{
		author.GET("/:id", getAuthor)
		author.GET("/all", getAuthors)
		author.POST("/add", addAuthor)
		author.PUT("/:id", replaceAuthor)
		author.PATCH("/:id", upgradeAuthor)
		author.DELETE("/:id", removeAuthor)
	}

	category := router.Group("/category")
	{
		category.GET("/:id", getCategory)
		category.GET("/all", getCategory)
		category.POST("/add", addCategory)
		category.PUT("/:id", replaceCategory)
		category.PATCH("/:id", upgradeCategory)
		category.DELETE("/:id", removeCategory)
	}

	//author routers
	return router
}
