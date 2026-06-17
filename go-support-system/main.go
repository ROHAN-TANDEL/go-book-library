package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

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
