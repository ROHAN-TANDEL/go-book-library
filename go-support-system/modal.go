package main

type Book struct {
	BookId          int `gorm:"column:book_id;primaryKey;autoIncrement" json:"book_id"`
	Title           string
	Language        string
	Publisher       string
	PublicationDate string `gorm:"column:publication_date" json:"publication_date"`
	Isbn            string `gorm:"column:isbn" json:"isbn"`
	Summary         string `gorm:"column:summary" json:"summary"`
}

type Author struct {
	AuthorID  int `gorm:"column:author_id;primaryKey;autoIncrement" json:"author_id"`
	Name      string
	Biography string
}

type Category struct {
	CategoryID  uint   `gorm:"column:category_id;primaryKey;autoIncrement" json:"category_id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
}
