package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {

	var dns string = "host=127.0.0.1 user=root password=root123 dbname=go_inventory sslmode=disable"
	db, err = gorm.Open(postgres.Open(dns))

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/get-book/:id", getBook)
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
	Summary         string
	Isbn            string
	Publisher       string
	PublicationDate string `gorm:"column:publication_date" json:"publication_date"`
}

func getBook(c *gin.Context) {

	var books []Book
	err = db.Find(&books).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	bookID := c.Param("id")
	c.JSON(200, gin.H{"action": "book is fetched", "book_id": bookID, "book": books})
}

type newBook struct {
	Title           string `json:"title"`
	Language        string `json:"language"`
	Summary         string `json:"summary"`
	Isbn            string `json:"isbn"`
	Publisher       string `json:"publisher"`
	PublicationDate string `json:"publication_date"`
}

func addBook(c *gin.Context) {
	var newBook newBook
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "message": "could not add the book"})
	}

	var record Book = Book{
		Title:           newBook.Title,
		Summary:         newBook.Summary,
		Isbn:            newBook.Isbn,
		Publisher:       newBook.Publisher,
		PublicationDate: newBook.PublicationDate,
		Language:        newBook.Language,
	}

	res := db.Create(&record)
	if res.Error != nil {
		c.JSON(500, gin.H{"error": res.Error})
	}

	c.JSON(200, gin.H{"action": "book is added", "book": newBook})
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
