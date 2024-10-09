package repo

import (
	"context"
	"database/sql"
	"library/internal/store/model"
	"library/internal/store/query"
	"math"
)

type BookRepo struct{}

var _ IBookRepo = BookRepo{}

func NewBookRepo() BookRepo {
	return BookRepo{}
}

func (r BookRepo) GetBookWithAuthors(ctx context.Context, bookID int64) (result model.BookWithAuthors, err error) {
	err = query.Q.Transaction(func(tx *query.Query) error {
		b := tx.Book
		a := tx.Author
		ba := tx.BookToAuthor

		book, err := b.WithContext(ctx).
			Where(b.ID.Eq(bookID)).
			First()
		if err != nil {
			return err
		}

		authors, err := a.WithContext(ctx).
			Join(ba, ba.AuthorID.EqCol(a.ID)).
			Where(ba.BookID.Eq(book.ID)).
			Find()
		if err != nil {
			return err
		}

		result = model.BookWithAuthors{
			Book:    book,
			Authors: authors,
		}
		return nil
	}, &sql.TxOptions{ReadOnly: true})
	return
}

const MaxPageLimit uint = 20

func (r BookRepo) GetBooksWithAuthors(ctx context.Context, page, limit uint) (results []model.BookWithAuthors, totalPages uint, err error) {
	err = query.Q.Transaction(func(tx *query.Query) error {
		b := tx.Book
		a := tx.Author
		ba := tx.BookToAuthor

		limit := int(min(MaxPageLimit, limit))
		books, _, err := b.WithContext(ctx).
			Order(b.ID.Asc()).
			FindByPage(max(0, int(page))*limit, limit)
		if err != nil {
			return err
		}

		totalBooks, err := b.WithContext(ctx).Count()
		if err != nil {
			return err
		}

		totalPages = uint(math.Ceil(float64(totalBooks) / float64(limit)))

		results = make([]model.BookWithAuthors, 0, len(books))
		for _, book := range books {
			var authors []*model.Author
			authors, err = a.WithContext(ctx).
				Join(ba, ba.AuthorID.EqCol(a.ID)).
				Where(ba.BookID.Eq(book.ID)).
				Find()
			if err != nil {
				return err
			}

			results = append(results, model.BookWithAuthors{
				Book:    book,
				Authors: authors,
			})
		}

		return nil
	}, &sql.TxOptions{ReadOnly: true})
	return
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
