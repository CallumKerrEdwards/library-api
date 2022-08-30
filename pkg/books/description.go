package books

import (
	"github.com/CallumKerrEdwards/library-api/pkg/books/text"
)

// Description - representation of a blurb of a book.
type Description struct {
	Text   string      `json:"text"`
	Format text.Format `json:"format,omitempty"`
}
