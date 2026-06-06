package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

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

type AuthorPatch struct {
	Name      *string
	Biography *string
}

type CategoryPatch struct {
	Name        *string
	Description *string
}

func main() {
	db = connect()
	var router *gin.Engine

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	router = route(router)
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
