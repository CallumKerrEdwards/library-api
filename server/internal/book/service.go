package book

import (
	"context"

	"github.com/CallumKerrEdwards/library/server/pkg/log"
)

type Store interface {
	GetBook(ctx context.Context, id string) (Book, error)
}

// Service - provides all functions for accessing and modifying Books
type Service struct {
	Store
	Log log.Logger
}

func NewService(store Store) *Service {
	return &Service{Store: store}
}

// GetBook - getting book
func (s *Service) GetBook(ctx context.Context, id string) (Book, error) {
	s.Log.WithField("id", id).Infof("getting book")

	book, err := s.Store.GetBook(ctx, id)
	if err != nil {
		s.Log.WithError(err).Debugln("Could not get book")
		return Book{}, err
	}
	return book, nil
}
