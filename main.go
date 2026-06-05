package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

type newBook struct {
	Title           *string `json:"title"`
	Language        *string `json:"language"`
	Summary         *string `json:"summary"`
	Isbn            *string `json:"isbn"`
	Publisher       *string `json:"publisher"`
	PublicationDate *string `json:"publication_date"`
}

func main() {
	db = connect()
	var router *gin.Engine

	router = route(router)
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
