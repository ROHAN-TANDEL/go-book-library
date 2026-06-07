package main

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
