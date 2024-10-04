package mock

import (
	"context"
	"github.com/google/uuid"
	"library/internal/store/model"

	"github.com/stretchr/testify/mock"
)

type UserStoreMock struct {
	mock.Mock
}

func (m *UserStoreMock) CreateUser(ctx context.Context, email string, password string) error {
	args := m.Called(ctx, email, password)

	return args.Error(0)
}

func (m *UserStoreMock) GetUser(ctx context.Context, email string) (*model.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*model.User), args.Error(1)
}

type SessionStoreMock struct {
	mock.Mock
}

func (m *SessionStoreMock) CreateSession(ctx context.Context, userId int64) (*model.Session, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).(*model.Session), args.Error(1)
}

func (m *SessionStoreMock) GetUserFromSession(ctx context.Context, sessionID uuid.UUID, userID int64) (*model.User, error) {
	args := m.Called(ctx, sessionID, userID)
	return args.Get(0).(*model.User), args.Error(1)
}
