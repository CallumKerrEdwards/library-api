package service

import "context"

// DeleteBook - delete Book.
func (s *Service) DeleteBook(ctx context.Context, id string) error {
	return s.Store.DeleteBook(ctx, id)
}
