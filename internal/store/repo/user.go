package repo

import (
	"context"
	"gorm.io/gorm"
	"library/internal/hash"
	"library/internal/store/model"
	"library/internal/store/query"
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

func (s *UserStore) CreateUser(ctx context.Context, email string, password string) error {

	hashedPassword, err := s.passwordHasher.GenerateFromPassword(password)
	if err != nil {
		return err
	}

	return query.Use(s.db).WithContext(ctx).User.Create(&model.User{
		Email:        email,
		PasswordHash: hashedPassword,
	})
}

func (s *UserStore) GetUser(ctx context.Context, email string) (*model.User, error) {
	u := query.User
	return query.Use(s.db).WithContext(ctx).User.Where(u.Email.Eq(email)).First()
}
