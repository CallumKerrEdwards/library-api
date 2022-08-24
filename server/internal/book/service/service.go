package service

import (
	"context"
	"errors"

	"github.com/CallumKerrEdwards/library/server/pkg/books"
	"github.com/CallumKerrEdwards/library/server/pkg/log"
)

var (
	ErrNotImplemented = errors.New("Not implemented")
)

type Store interface {
	GetBook(ctx context.Context, id string) (books.Book, error)
	GetAllBooks(ctx context.Context) ([]books.Book, error)
	PostBook(ctx context.Context, book books.Book) (books.Book, error)
}

// Service - provides all functions for accessing and modifying Books
type Service struct {
	Store
	Log log.Logger
}

func NewService(store Store, logger log.Logger) *Service {
	return &Service{Store: store, Log: logger}
}
