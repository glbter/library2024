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

type NewUserRepoParams struct {
	PasswordHasher hash.PasswordHasher
}

func NewUserRepo(params NewUserRepoParams) *UserRepo {
	return &UserRepo{
		passwordHasher: params.PasswordHasher,
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, email string, password string) error {

	hashedPassword, err := r.passwordHasher.GenerateFromPassword(password)
	if err != nil {
		return err
	}

	return query.User.WithContext(ctx).Create(&model.User{
		Email:        email,
		PasswordHash: hashedPassword,
	})
}

func (r *UserRepo) GetUser(ctx context.Context, email string) (*model.User, error) {
	u := query.User
	return u.WithContext(ctx).
		Where(u.Email.Eq(email)).
		Take()
}
