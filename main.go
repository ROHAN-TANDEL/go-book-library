package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	/*
		get:
		curl -X GET 'http://localhost:8080/get-books'

		get:
		curl -X GET 'http://localhost:8080/get-book/28'

		put:
		curl -X PUT "http://localhost:8080/replace-book/19" -d '{"title": "RAHi 2", "language": "Arabic", "summary": "Sparrow and the rain", "isbn": "122222WRRHAN", "publisher": "Trains and Tracks Co", "publication_date":"2000-10-10"}' -H "Content-Type: application/json"

		post:
		curl -X POST "http://localhost:8080/add-book" -d '{"title": "The Power Of Subconcious Mind", "language": "Arabic", "summary": "Sparrow and the rain", "isbn": "889a891244403E", "publisher": "Trains and Tracks Co", "publication_date":"2000-10-10"}' -H "Content-Type: application/json"

		patch:
		curl -X PATCH "http://localhost:8080/upgrade-book/28" -d '{"title": "Power of Habits", "language": "Arabic", "summary": "Sparrow and the rain", "isbn": "122222WWWARR", "publisher": "Trains and Tracks Co", "publication_date":"2000-10-10"}' -H "Content-Type: application/json"

		delete:
		curl -X DELETE "http://localhost:8080/remove-book/28"
	*/
	var err error
	var dns string = "host=127.0.0.1 user=root password=root123 dbname=go_inventory sslmode=disable"
	db, err = gorm.Open(postgres.Open(dns))

	if err != nil {
		panic(err)
	}

	var router = gin.Default()

	router.GET("/get-book/:id", getBook)
	router.GET("/get-books", getBooks)
	router.POST("/add-book", addBook)
	router.PUT("/replace-book/:id", replaceBook)
	router.PATCH("/upgrade-book/:id", upgradeBook)
	router.DELETE("/remove-book/:id", removeBook)
	router.Run(":8080")
}

type Book struct {
	BookId          int `gorm:"column:book_id;primaryKey;AutoIncrement" json:"book_id"`
	Title           string
	Language        string
	Publisher       string
	PublicationDate string `gorm:"column:publication_date" json:"publication_date"`
	Isbn            string `gorm:"column:isbn" json:"isbn"`
	Summary         string `gorm:"column:summary" json:"summary"`
}

func getBooks(c *gin.Context) {
	var books []Book

	if err := db.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func getBook(c *gin.Context) {
	var books Book
	var bookId = c.Param("id")

	err := db.Where("book_id = ?", bookId).First(&books)
	if err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error.Error()})
		return
	}

	if err.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

type newBook struct {
	Title           *string `json:"title"`
	Language        *string `json:"language"`
	Summary         *string `json:"summary"`
	Isbn            *string `json:"isbn"`
	Publisher       *string `json:"publisher"`
	PublicationDate *string `json:"publication_date"`
}

func addBook(c *gin.Context) {
	var record newBook
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var data = Book{
		Title:           *record.Title,
		Publisher:       *record.Publisher,
		Language:        *record.Language,
		Isbn:            *record.Isbn,
		PublicationDate: *record.PublicationDate,
		Summary:         *record.Summary,
	}

	if err := db.Create(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func removeBook(c *gin.Context) {
	var bookID = c.Param("id")

	err := db.Where("book_id = ?", bookID).Delete(&Book{})

	if err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error.Error()})
		return
	}

	if err.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Book does not exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookID})
}

func replaceBook(c *gin.Context) {
	var record newBook
	var bookID = c.Param("id")
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row := db.Model(&Book{}).Where("book_id = ?", bookID).Updates(&record)
	if row.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": row.Error.Error()})
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

	var update = make(map[string]interface{})

	if record.Title != nil {
		update["title"] = record.Title
	}

	if record.Language != nil {
		update["language"] = record.Language
	}

	if record.Summary != nil {
		update["summary"] = record.Summary
	}

	if record.Isbn != nil {
		update["isbn"] = record.Isbn
	}

	if record.Publisher != nil {
		update["publisher"] = record.Publisher
	}

	if record.PublicationDate != nil {
		update["publication_date"] = record.PublicationDate
	}

	res := db.Model(&Book{}).Where("book_id = ?", bookId).Updates(&update)

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		return
	}

	if res.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Book does not exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": record})
}
