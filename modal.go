package main

type Book struct {
	BookId          int `gorm:"column:book_id;primaryKey;AutoIncrement" json:"book_id"`
	Title           string
	Language        string
	Publisher       string
	PublicationDate string `gorm:"column:publication_date" json:"publication_date"`
	Isbn            string `gorm:"column:isbn" json:"isbn"`
	Summary         string `gorm:"column:summary" json:"summary"`
}

type Author struct {
	Author    int `gorm:"column:author_id;PrimaryKey;AutoIncrement" json:"author_id"`
	Name      string
	Biography string
}

type Category struct {
	CategoryId  uint   `gorm:"column:category_id;PrimaryKey;AutoIncrement" json:"category_id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
}
