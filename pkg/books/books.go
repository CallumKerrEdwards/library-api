package books

import "github.com/google/uuid"

// Book - representation of a book
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Series Series `json:"series"`
}

// Series - representation of a series of books
type Series struct {
	Sequence int    `json:"sequence"`
	Title    string `json:"title"`
}

func NewBook(title, author string, series Series) Book {
	return Book{
		ID:     uuid.New().String(),
		Title:  title,
		Author: author,
		Series: series,
	}
}
