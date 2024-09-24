package dbstore

import (
	"library/internal/hash"
	"library/internal/store"

	"gorm.io/gorm"
)

type UserStore struct {
	db             *gorm.DB
	passwordHasher hash.PasswordHasher
}

type NewUserStoreParams struct {
	DB             *gorm.DB
	PasswordHasher hash.PasswordHasher
}

func NewUserStore(params NewUserStoreParams) *UserStore {
	return &UserStore{
		db:             params.DB,
		passwordHasher: params.PasswordHasher,
	}
}

func (s *UserStore) CreateUser(email string, password string) error {

	hashedPassword, err := s.passwordHasher.GenerateFromPassword(password)
	if err != nil {
		return err
	}

	return s.db.Create(&store.User{
		Email:        email,
		PasswordHash: hashedPassword,
	}).Error
}

func (s *UserStore) GetUser(email string) (*store.User, error) {

	var user store.User
	err := s.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, err
}
