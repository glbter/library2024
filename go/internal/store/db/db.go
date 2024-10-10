package db

import (
	"library/internal/store/model"
	"library/internal/store/query"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func open(dsn string) (*gorm.DB, error) {
	// make the temp directory if it doesn't exist
	err := os.MkdirAll("/tmp", 0755)
	if err != nil {
		return nil, err
	}

	return gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		//PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
}

func MustOpen(dsn string) *gorm.DB {
	db, err := open(dsn)
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Session{},
		&model.Book{},
		&model.Genre{},
		&model.BookToGenre{},
		&model.Author{},
		&model.BookToAuthor{},
		&model.BookLendRequest{},
		&model.BookLendTransaction{},
		&model.BookReturnTransaction{},
	)
	if err != nil {
		panic(err)
	}

	query.SetDefault(db)

	return db
}
