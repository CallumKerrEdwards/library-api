//go:build integration
// +build integration

package db

import (
	"context"
	"testing"

	"github.com/CallumKerrEdwards/loggerrific/tlogger"
	"github.com/stretchr/testify/assert"

	"github.com/CallumKerrEdwards/library-api/pkg/books"
)

func TestBookDatabase(t *testing.T) {
	t.Run("test create book", func(t *testing.T) {
		// given
		db, err := NewDatabase(context.Background(), tlogger.NewTLogger(t))
		assert.NoError(t, err)
		db.clearBooksCollection(context.Background())

		book, err := db.PostBook(context.Background(), &books.Book{
			ID:      "001",
			Title:   "A Wizard of Earthsea",
			Authors: []books.Person{{Forenames: "Ursula K", SortName: "Le Guin"}},
		})
		assert.NoError(t, err)

		// when
		newBook, err := db.GetBook(context.Background(), book.ID)

		// then
		assert.NoError(t, err)
		assert.Equal(t, "Le Guin", newBook.Authors[0].SortName)
	})

	t.Run("test delete book", func(t *testing.T) {
		// given
		db, err := NewDatabase(context.Background(), tlogger.NewTLogger(t))
		assert.NoError(t, err)
		err = db.clearBooksCollection(context.Background())
		assert.NoError(t, err)

		book, err := db.PostBook(context.Background(), &books.Book{
			ID:      "002",
			Title:   "The Tombs of Atuan",
			Authors: []books.Person{{Forenames: "Ursula K", SortName: "Le Guin"}},
		})
		assert.NoError(t, err)

		// when
		err = db.DeleteBook(context.Background(), book.ID)

		// then
		assert.NoError(t, err)
		_, err = db.GetBook(context.Background(), book.ID)
		assert.Error(t, err)
	})
}

func TestBookDatabase_GetAudiobooksOnly(t *testing.T) {
	t.Run("test get only books with audiobooks", func(t *testing.T) {
		// given
		db, err := NewDatabase(context.Background(), tlogger.NewTLogger(t))
		assert.NoError(t, err)
		db.clearBooksCollection(context.Background())

		_, err = db.PostBook(context.Background(), &books.Book{
			ID:      "001",
			Title:   "A Wizard of Earthsea",
			Authors: []books.Person{{Forenames: "Ursula K", SortName: "Le Guin"}},
			Audiobook: &books.Audiobook{
				Narrators: []books.Person{{Forenames: "Kobna", SortName: "Holdbrook-Smith"}},
			},
		})
		assert.NoError(t, err)
		_, err = db.PostBook(context.Background(), &books.Book{
			ID:      "002",
			Title:   "The Tombs of Atuan",
			Authors: []books.Person{{Forenames: "Ursula K", SortName: "Le Guin"}},
		})
		assert.NoError(t, err)

		// when
		allAudiobooks, err := db.GetAllBooksWithAudiobook(context.Background())

		// then
		assert.NoError(t, err)
		assert.Len(t, allAudiobooks, 1)
		assert.Equal(t, "Holdbrook-Smith", allAudiobooks[0].Audiobook.Narrators[0].SortName)
	})
}
