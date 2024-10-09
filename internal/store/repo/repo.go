package repo

import (
	"context"
	"github.com/google/uuid"
	"library/internal/store/model"
)

type IUserRepo interface {
	CreateUser(ctx context.Context, email string, password string) error
	GetUser(ctx context.Context, email string) (*model.User, error)
}

type ISessionRepo interface {
	CreateSession(ctx context.Context, userId int64) (*model.Session, error)
	GetUserFromSession(ctx context.Context, sessionID uuid.UUID, userID int64) (*model.User, error)
}

type IBookRepo interface {
	GetBooksWithAuthors(ctx context.Context, page, limit uint) (books []model.BookWithAuthors, totalPages uint, err error)
	GetBookWithAuthors(ctx context.Context, bookID int64) (model.BookWithAuthors, error)
	RequestBook(ctx context.Context, userID, bookID int64) error
}

type IAuthorRepo interface {
	GetAuthorWithBooks(ctx context.Context, authorID int64) (model.AuthorWithBooks, error)
}
