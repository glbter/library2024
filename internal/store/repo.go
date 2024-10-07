package store

import (
	"context"
	"github.com/google/uuid"
	"library/internal/store/model"
)

type UserRepo interface {
	CreateUser(ctx context.Context, email string, password string) error
	GetUser(ctx context.Context, email string) (*model.User, error)
}

type SessionRepo interface {
	CreateSession(ctx context.Context, userId int64) (*model.Session, error)
	GetUserFromSession(ctx context.Context, sessionID uuid.UUID, userID int64) (*model.User, error)
}

type BookRepo interface {
	GetBooksWithAuthors(ctx context.Context, page, limit uint) (books []BookWithAuthors, totalPages uint, err error)
	RequestBook(ctx context.Context, UserID, BookID int64) error
}

type BookWithAuthors struct {
	Book    *model.Book
	Authors []*model.Author
}
