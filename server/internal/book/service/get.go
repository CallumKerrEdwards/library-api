package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/CallumKerrEdwards/library/server/pkg/books"
)

var (
	ErrFetchingBook = errors.New("failed to fetch Book by id")
)

// GetBook - getting book
func (s *Service) GetBook(ctx context.Context, id string) (books.Book, error) {
	s.Log.WithField("id", id).Infof("getting Book")

	book, err := s.Store.GetBook(ctx, id)
	if err != nil {
		s.Log.WithError(err).Errorln("Get Failed")
		return books.Book{}, fmt.Errorf("%w %s", ErrFetchingBook, id)
	}
	return book, nil
}
