package book

import (
	"context"
	"errors"

	"github.com/CallumKerrEdwards/library/server/pkg/log"
)

var (
	ErrNotImplemented = errors.New("Not implemented")
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
