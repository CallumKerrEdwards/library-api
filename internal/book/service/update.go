package service

import (
	"context"

	"github.com/CallumKerrEdwards/library-api/pkg/books"
)

// UpdateBook - update book details.
func (s *Service) UpdateBook(
	ctx context.Context,
	id string,
	updatedBook *books.Book,
) (books.Book, error) {
	updated, newlyUpdatedBook, err := s.Store.UpdateBook(ctx, id, updatedBook)
	if err != nil {
		s.Log.WithError(err).Errorln("Error updating book with ID", id)
		return books.Book{}, err
	}

	if !updated {
		s.Log.Infoln("No book found with ID", id, "to update")
	}

	return newlyUpdatedBook, nil
}
