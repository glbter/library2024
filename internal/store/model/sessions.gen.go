// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/google/uuid"

const TableNameSession = "sessions"

// Session mapped from table <sessions>
type Session struct {
	ID     uuid.UUID `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID int64     `gorm:"column:user_id;type:bigint;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user_id"`
}

// TableName Session's table name
func (*Session) TableName() string {
	return TableNameSession
}
