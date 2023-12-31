package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrCreatingComment = errors.New("failed to create comment")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment - a representation of the comment
// structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, Comment) (Comment, error)
}

// Service - is the struct on which all our
// logic will be built on top of
type Service struct {
	Store
}

// NewService - returns pointer to new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// here we accept interface in param and return
// struct from func, good architecture for testing

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, ID string, updatedCmt Comment) (Comment, error) {
	fmt.Println("updating a comment")
	cmt, err := s.Store.UpdateComment(ctx, ID, updatedCmt)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}

	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	fmt.Println("deleting a comment")
	err := s.Store.DeleteComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	fmt.Println("posting a comment")
	insertedCmt, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrCreatingComment
	}

	return insertedCmt, nil
}
