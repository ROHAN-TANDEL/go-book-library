package main

import "github.com/gin-gonic/gin"

func route(router *gin.Engine) *gin.Engine {

	router = gin.Default()

	router.GET("/get-book/:id", getBook)
	router.GET("/get-books", getBooks)
	router.POST("/add-book", addBook)
	router.PUT("/replace-book/:id", replaceBook)
	router.PATCH("/upgrade-book/:id", upgradeBook)
	router.DELETE("/remove-book/:id", removeBook)

	return router
}
