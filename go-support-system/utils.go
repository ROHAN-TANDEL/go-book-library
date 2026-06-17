package main

type newBook struct {
	BookID          int     `gorm:"column:book_id;primaryKey;autoIncrement" json:"book_id"`
	Title           *string `json:"title"`
	Language        *string `json:"language"`
	Summary         *string `json:"summary"`
	Isbn            *string `json:"isbn"`
	Publisher       *string `json:"publisher"`
	PublicationDate *string `json:"publication_date"`

	// GORM automatic relationship declarations
	Authors    []Author   `gorm:"many2many:book_authors;foreignKey:BookID;joinForeignKey:book_id;References:AuthorID;joinReferences:author_id"`
	Categories []Category `gorm:"many2many:book_categories;foreignKey:BookID;joinForeignKey:book_id;References:CategoryID;joinReferences:category_id"`
	//	Authors    []Author   `gorm:"many2many:book_author;"`
	//	Categories []Category `gorm:"many2many:book_categories;"`
}

func (newBook) TableName() string {
	return "books"
}

type AuthorPatch struct {
	Name      *string
	Biography *string
}

type CategoryPatch struct {
	Name        *string
	Description *string
}
