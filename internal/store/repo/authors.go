package repo

import (
	"context"
	"database/sql"
	"library/internal/store/model"
	"library/internal/store/query"
)

type AuthorRepo struct{}

var _ IAuthorRepo = AuthorRepo{}

func NewAuthorRepo() AuthorRepo {
	return AuthorRepo{}
}

func (r AuthorRepo) GetAuthorWithBooks(ctx context.Context, authorID int64) (model.AuthorWithBooks, error) {
	var result model.AuthorWithBooks
	err := query.Q.Transaction(func(tx *query.Query) error {
		b := tx.Book
		a := tx.Author
		ba := tx.BookToAuthor

		author, err := a.WithContext(ctx).
			Where(a.ID.Eq(authorID)).
			First()
		if err != nil {
			return err
		}

		books, err := b.WithContext(ctx).
			Join(ba, ba.BookID.EqCol(b.ID)).
			Where(ba.AuthorID.Eq(authorID)).
			Order(b.PublishedOn.Desc()).
			Find()
		if err != nil {
			return err
		}

		result = model.AuthorWithBooks{
			Books:  books,
			Author: author,
		}
		return nil
	}, &sql.TxOptions{ReadOnly: true})

	return result, err
}
