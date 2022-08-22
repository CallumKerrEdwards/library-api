package service

import (
	"context"

	"github.com/CallumKerrEdwards/library/server/pkg/books"
)

// CreateBook - create new book
func (s *Service) PostBook(ctx context.Context, book books.Book) (books.Book, error) {
	s.Log.WithField("id", book.ID).Infoln("creating Book")

	createdBook, err := s.Store.PostBook(ctx, book)
	if err != nil {
		s.Log.WithError(err).Errorln("Get Failed")
		return books.Book{}, err
	}
	return createdBook, nil
}
