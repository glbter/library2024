package store

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" gorm:"uniqueIndex"`
	PasswordHash string `json:"-"`

	BookLendRequests []BookLendRequest `gorm:"many2many:book_lend_requests;"`
}

type Session struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	SessionID string `json:"session_id"`
	UserID    uint   `json:"user_id"`
	User      User   `gorm:"foreignKey:UserID" json:"user"`
}

type Book struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	PublishedAt time.Time
	Amount      uint
	Authors     []*Author `gorm:"many2many:book_authors;"`
	Genres      []*Genre  `gorm:"many2many:book_genres;"`

	LendRequests       []BookLendRequest       `gorm:"many2many:book_lend_requests;"`
	LendTransactions   []BookLendTransaction   `gorm:"many2many:book_lend_transactions;"`
	ReturnTransactions []BookReturnTransaction `gorm:"many2many:book_return_transactions;"`
}

type Author struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string

	Books []*Book `gorm:"many2many:book_authors;"`
}

type BookLendTransaction struct {
	gorm.Model
	UserID       uint
	BookID       uint
	Request      BookLendRequest
	CreatedAt    time.Time
	DateToReturn time.Time
}

type BookReturnTransaction struct {
	gorm.Model
	UserID          uint
	BookID          uint
	LendTransaction BookLendTransaction
	CreatedAt       time.Time
}

type BookLendRequest struct {
	gorm.Model
	UserID uint
	BookID uint
}

type Genre struct {
	gorm.Model
	ID   uint
	Name string

	Books []*Book `gorm:"many2many:book_genres;"`
}

type UserStore interface {
	CreateUser(email string, password string) error
	GetUser(email string) (*User, error)
}

type SessionStore interface {
	CreateSession(session *Session) (*Session, error)
	GetUserFromSession(sessionID string, userID string) (*User, error)
}
