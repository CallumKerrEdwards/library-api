package books

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/CallumKerrEdwards/library-api/pkg/books/genres"
)

// Book - representation of a book.
type Book struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Authors     []Person       `json:"authors"`
	Description *Description   `json:"description,omitempty"`
	ReleaseDate *ReleaseDate   `json:"releaseDate,omitempty"`
	Genres      []genres.Genre `json:"genres,omitempty"`
	Series      Series         `json:"series"`
	Audiobook   *Audiobook     `json:"audiobook,omitempty"`
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

func NewBook(title string, description *Description, authors []Person, releaseDate *ReleaseDate,
	genreList []genres.Genre, series Series, audiobook *Audiobook) Book {
	return Book{
		ID:          uuid.New().String(),
		Title:       title,
		Authors:     authors,
		Description: description,
		ReleaseDate: releaseDate,
		Genres:      genreList,
		Series:      series,
		Audiobook:   audiobook,
	}
}

func (b *Book) GetAuthor() string {
	return GetPersonsString(b.Authors)
}

func GetPersonsString(p []Person) string {
	switch len(p) {
	case 0:
		return ""
	case 1:
		return p[0].GetPersonString()
	default:
		var personStrs []string
		for _, person := range p {
			personStrs = append(personStrs, person.GetPersonString())
		}

		return fmt.Sprintf("%s & %s", strings.Join(personStrs[:len(personStrs)-1], ", "), personStrs[len(personStrs)-1])
	}
}

func (p Person) GetPersonString() string {
	return fmt.Sprintf("%s %s", p.Forenames, p.SortName)
}
