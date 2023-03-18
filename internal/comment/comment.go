package comment

import (
	"context"
	"errors"
)

var (
	ErrFetchingComment = errors.New("no comment was found")
	ErrNotImplemented  = errors.New("not implemented")
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	comment, err := s.Store.GetComment(ctx, id)
	if err != nil {
		return Comment{}, ErrFetchingComment
	}
	return comment, nil
}

func (s *Service) CreateComment(ctx context.Context, comment Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}

func (s *Service) UpdateComment(ctx context.Context, comment Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, comment Comment) error {
	return ErrNotImplemented
}
