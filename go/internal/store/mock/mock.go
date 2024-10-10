package mock

import (
	"context"
	"github.com/google/uuid"
	"library/internal/store/model"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (m *UserRepoMock) CreateUser(ctx context.Context, email string, password string) error {
	args := m.Called(ctx, email, password)

	return args.Error(0)
}

func (m *UserRepoMock) GetUser(ctx context.Context, email string) (*model.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*model.User), args.Error(1)
}

type SessionRepoMock struct {
	mock.Mock
}

func (m *SessionRepoMock) CreateSession(ctx context.Context, userId int64) (*model.Session, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).(*model.Session), args.Error(1)
}

func (m *SessionRepoMock) GetUserFromSession(ctx context.Context, sessionID uuid.UUID, userID int64) (*model.User, error) {
	args := m.Called(ctx, sessionID, userID)
	return args.Get(0).(*model.User), args.Error(1)
}
