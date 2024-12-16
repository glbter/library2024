package repo

import (
	"context"
	"library/internal/hash"
	"library/internal/store/model"
	"library/internal/store/query"
)

type UserRepo struct {
	passwordHasher hash.PasswordHasher
}

var _ IUserRepo = &UserRepo{}

func NewUserRepo(passwordHasher hash.PasswordHasher) *UserRepo {
	return &UserRepo{passwordHasher}
}

func (r *UserRepo) CreateUser(ctx context.Context, email string, password string) error {
	hashedPassword, err := r.passwordHasher.GenerateFromPassword(password)
	if err != nil {
		return err
	}

	return query.Q.Transaction(func(tx *query.Query) error {
		maxIdUser, err := tx.User.WithContext(ctx).
			Select(tx.User.ID).
			Order(tx.User.ID.Desc()).
			Limit(1).
			First()
		if err != nil {
			return err
		}

		return tx.User.WithContext(ctx).Create(&model.User{
			ID:           maxIdUser.ID + 1,
			Email:        email,
			PasswordHash: hashedPassword,
		})
	})
}

func (r *UserRepo) GetUser(ctx context.Context, email string) (*model.User, error) {
	u := query.User
	return u.WithContext(ctx).
		Where(u.Email.Eq(email)).
		Take()
}
