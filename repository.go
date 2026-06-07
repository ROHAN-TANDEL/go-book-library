package main

import "gorm.io/gorm"

func addBookRepo(record *newBook) (Book, error) {
	var data = Book{
		Title:           *record.Title,
		Publisher:       *record.Publisher,
		Language:        *record.Language,
		Isbn:            *record.Isbn,
		PublicationDate: *record.PublicationDate,
		Summary:         *record.Summary,
	}

	err := db.Create(&data)

	return data, err.Error

}

func getBookRepo(bookId any) (Book, error) {
	var books Book
	err := db.Where("book_id = ?", bookId).First(&books)
	return books, err.Error
}

func getBooksRepo(page map[string]int) ([]Book, error) {
	var books []Book
	booksList := paginateRepo(db, page)
	res := booksList.Find(&books)
	return books, res.Error
}

func removeBookRepo(bookId any) (int64, error) {
	res := db.Where("book_id = ?", bookId).Delete(&Book{})
	return res.RowsAffected, res.Error
}

func replaceBookRepo(bookId any, record newBook) (newBook, error) {
	row := db.Model(&Book{}).Where("book_id = ?", bookId).Updates(&record)
	return record, row.Error
}

func upgradeBookRepo(bookId any, record newBook) (int64, error) {

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
	return res.RowsAffected, res.Error
}

func getAuthorRepo(authorId any) (Author, error) {
	var author Author
	res := db.Where("author_id = ?", authorId).First(&author)
	return author, res.Error
}

func getAuthorsRepo() ([]Author, error) {
	var authors []Author
	res := db.Find(&authors)
	return authors, res.Error
}

func addAuthorRepo(newAuthorInput Author) (Author, error) {
	record := Author{
		Name:      newAuthorInput.Name,
		Biography: newAuthorInput.Biography,
	}
	res := db.Create(&record)

	return record, res.Error
}

func replaceAuthorRepo(authorId any, record Author) (Author, error) {
	res := db.Model(&Author{}).Where("author_id = ?", authorId).Updates(&record)
	return record, res.Error
}

func upgradeAuthorRepo(authorId any, record AuthorPatch) (AuthorPatch, error) {
	if authorId == nil {
		return AuthorPatch{}, nil
	}
	var update = make(map[string]interface{})

	if record.Name != nil {
		update["name"] = record.Name
	}

	if record.Biography != nil {
		update["biography"] = record.Biography
	}

	if len(update) != 0 {
		res := db.Model(&Author{}).Where("author_id = ?", authorId).Updates(&update)
		return record, res.Error
	}

	return AuthorPatch{}, nil
}

func removeAuthorRepo(authorId any) (int64, error) {
	res := db.Where("author_id = ?", authorId).Delete(&Author{})
	return res.RowsAffected, res.Error
}

func getCategoryRepo(catID any) (Category, error) {
	var category Category
	err := db.Where("category_id = ?", catID).First(&category)
	return category, err.Error
}

func getCategoriesRepo(page map[string]int) ([]Category, error) {
	var categories []Category
	db = paginateRepo(db, page)
	data := db.Debug().Find(&categories)
	return categories, data.Error
}

func addCategoryRepo(category Category) (Category, error) {
	var record = Category{
		Name:        category.Name,
		Description: category.Description,
	}

	res := db.Create(&record)

	return record, res.Error
}

func replaceCategoryRepo(catID any, category Category) (Category, error) {
	var record = Category{
		Name:        category.Name,
		Description: category.Description,
	}
	res := db.Where("category_id = ?", catID).Updates(&record)

	return record, res.Error
}

func upgradeCategoryRepo(catID any, category CategoryPatch) (CategoryPatch, error) {

	var record = make(map[string]any)

	if category.Name != nil {
		record["name"] = category.Name
	}

	if category.Description != nil {
		record["description"] = category.Description
	}

	if len(record) == 0 {
		return CategoryPatch{}, nil
	}

	res := db.Model(&Category{}).Where("category_id = ?", catID).Updates(&record)

	return category, res.Error
}

func removeCategoryRepo(catID any) (any, error) {
	res := db.Where("category_id = ?", catID).Delete(&Category{})
	return catID, res.Error
}

func paginateRepo(db *gorm.DB, filters map[string]int) *gorm.DB {
	var tempdb = db
	if filters != nil {
		if filters["limit"] > 0 {
			tempdb = tempdb.Limit(filters["limit"])
		} else {
			tempdb = tempdb.Limit(2)
		}

		if filters["offset"] > 0 {
			tempdb = tempdb.Offset(filters["offset"])
		}
	}
	return tempdb
}
