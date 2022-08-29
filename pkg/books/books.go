package books

import (
	"github.com/google/uuid"
)

// Book - representation of a book.
type Book struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Authors     []Person     `json:"authors"`
	Description string       `json:"description,omitempty"`
	ReleaseDate *ReleaseDate `json:"releaseDate,omitempty"`
	Genres      []Genre      `json:"genres,omitempty"`
	Series      Series       `json:"series"`
	Arefacts    []Artefact   `json:"artefacts,omitempty"`
}

// Person - represetation of a person, for example an author or audiobook narrator.
type Person struct {
	Forenames string `json:"forenames"`
	SortName  string `json:"sortName"`
}

// Series - representation of a series of books.
type Series struct {
	Sequence int    `json:"sequence"`
	Title    string `json:"title"`
}

type Artefact interface {
	GetPath() string
}

func NewBook(title, description string, authors []Person, releaseDate *ReleaseDate,
	genres []Genre, series Series, artefacts []Artefact) Book {
	return Book{
		ID:          uuid.New().String(),
		Title:       title,
		Authors:     authors,
		Description: description,
		ReleaseDate: releaseDate,
		Genres:      genres,
		Series:      series,
		Arefacts:    artefacts,
	}
}
