package main

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

func getBooksRepo() ([]Book, error) {
	var books []Book
	err := db.Find(&books)
	return books, err.Error
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
