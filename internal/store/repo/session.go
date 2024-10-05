package repo

import (
	"context"
	"library/internal/store/model"
	"library/internal/store/query"

	"github.com/google/uuid"
)

type SessionRepo struct{}

type NewSessionStoreParams struct{}

func NewSessionRepo() SessionRepo {
	return SessionRepo{}
}

func (r SessionRepo) CreateSession(ctx context.Context, userId int64) (*model.Session, error) {
	session := model.Session{UserID: userId}

	err := query.Session.WithContext(ctx).Create(&session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r SessionRepo) GetUserFromSession(ctx context.Context, sessionID uuid.UUID, userID int64) (*model.User, error) {
	u := query.User
	s := query.Session
	return u.WithContext(ctx).
		Join(s, u.ID.EqCol(s.UserID)).
		Where(s.ID.Eq(sessionID), u.ID.Eq(userID)).
		First()
}
