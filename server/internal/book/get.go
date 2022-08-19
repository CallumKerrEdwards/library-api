package book

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingBook = errors.New("failed to fetch Book by id")
)

// GetBook - getting book
func (s *Service) GetBook(ctx context.Context, id string) (Book, error) {
	s.Log.WithField("id", id).Infof("getting Book")

	book, err := s.Store.GetBook(ctx, id)
	if err != nil {
		s.Log.WithError(err).Errorln("Get Failed")
		return Book{}, fmt.Errorf("%w %s", ErrFetchingBook, id)
	}
	return book, nil
}
