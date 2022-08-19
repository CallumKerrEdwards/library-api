package book

import "context"

// CreateBook - create new book
func (s *Service) CreateBook(ctx context.Context, book Book) (Book, error) {
	return Book{}, ErrNotImplemented
}
