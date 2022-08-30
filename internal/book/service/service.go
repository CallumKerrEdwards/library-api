package service

import (
	"context"

	"github.com/CallumKerrEdwards/loggerrific"

	"github.com/CallumKerrEdwards/library-api/pkg/books"
)

type Store interface {
	GetBook(ctx context.Context, id string) (books.Book, error)
	GetAllBooks(ctx context.Context) ([]books.Book, error)
	PostBook(ctx context.Context, book *books.Book) (books.Book, error)
	DeleteBook(ctx context.Context, id string) error
	UpdateBook(ctx context.Context, id string, book *books.Book) (bool, books.Book, error)
}

// Service - provides all functions for accessing and modifying Books.
type Service struct {
	Store
	Log loggerrific.Logger
}

func NewService(store Store, logger loggerrific.Logger) *Service {
	return &Service{Store: store, Log: logger}
}
