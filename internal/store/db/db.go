package db

import (
	"library/internal/store"
	"os"

	"gorm.io/driver/postgres" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

func open(dsn string) (*gorm.DB, error) {

	// make the temp directory if it doesn't exist
	err := os.MkdirAll("/tmp", 0755)
	if err != nil {
		return nil, err
	}

	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

}

func MustOpen(dsn string) *gorm.DB {
	if dsn == "" {
		//TODO: [Hlib] insert local credentials for development purposes
		dsn = "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	}

	db, err := open(dsn)
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&store.User{}, &store.Session{})

	if err != nil {
		panic(err)
	}

	return db
}
