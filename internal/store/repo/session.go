package repo

import (
	"context"
	"library/internal/store/model"
	"library/internal/store/query"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SessionRepo struct {
	db *gorm.DB
}

type NewSessionStoreParams struct {
	DB *gorm.DB
}

func NewSessionStore(params NewSessionStoreParams) *SessionRepo {
	return &SessionRepo{
		db: params.DB,
	}
}

func (s *SessionRepo) CreateSession(ctx context.Context, userId int64) (*model.Session, error) {
	session := model.Session{UserID: userId}

	err := query.Use(s.db).WithContext(ctx).Session.Create(&session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *SessionRepo) GetUserFromSession(ctx context.Context, sessionID uuid.UUID, userID int64) (*model.User, error) {
	return query.Use(s.db).WithContext(ctx).
		User.
		Join(query.Session, query.User.ID.EqCol(query.Session.UserID)).
		Where(query.Session.ID.Eq(sessionID), query.User.ID.Eq(userID)).
		First()
}
