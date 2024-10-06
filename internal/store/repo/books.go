package repo

import (
	"context"
	"database/sql"
	"library/internal/store"
	"library/internal/store/model"
	"library/internal/store/query"
)

type BookRepo struct{}

func NewBookRepo() BookRepo {
	return BookRepo{}
}

type ID int64

func (r BookRepo) GetBooksWithAuthors(ctx context.Context, page, limit int) ([]store.BookWithAuthors, error) {
	var results []store.BookWithAuthors
	err := query.Q.Transaction(func(tx *query.Query) error {
		b := tx.Book
		a := tx.Author
		ba := tx.BookToAuthor

		books, err := b.WithContext(ctx).
			Order(b.ID.Desc()).
			Limit(limit).
			Offset((page - 1) * limit).
			Find()
		if err != nil {
			return err
		}

		results = make([]store.BookWithAuthors, 0, len(books))
		for _, book := range books {
			var authors []*model.Author
			authors, err = a.WithContext(ctx).
				Join(ba, ba.AuthorID.EqCol(a.ID)).
				Where(ba.BookID.Eq(book.ID)).
				Find()
			if err != nil {
				return err
			}

			results = append(results, store.BookWithAuthors{
				Book:    book,
				Authors: authors,
			})
		}

		return nil
	}, &sql.TxOptions{ReadOnly: true})
	return results, err
}

func (r BookRepo) RequestBook(ctx context.Context, userID, bookID int64) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		lr := tx.BookLendRequest

		bookLendRequest := model.BookLendRequest{UserID: userID, BookID: bookID}
		if err := lr.WithContext(ctx).Create(&bookLendRequest); err != nil {
			return err
		}

		lt := tx.BookLendTransaction
		rt := tx.BookReturnTransaction

		// amount of taken books
		count, err := lt.WithContext(ctx).
			LeftJoin(rt, lr.ID.EqCol(rt.RequestID)).
			Where(lr.BookID.Eq(bookID), rt.RequestID.IsNull()).
			Count()
		if err != nil {
			return err
		}

		book := tx.Book
		dbBook, err := book.WithContext(ctx).
			Where(book.ID.EqCol(lt.RequestID)).
			First()
		if err != nil {
			return err
		}

		// lend a book
		if dbBook != nil && dbBook.Amount > int16(count) {
			err = lt.WithContext(ctx).
				Create(&model.BookLendTransaction{RequestID: bookLendRequest.ID.Bytes})
			if err != nil {
				return err
			}
		}

		return nil
	})
}
