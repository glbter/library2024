// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID           int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	FirstName    string  `gorm:"column:first_name;not null" json:"first_name"`
	LastName     *string `gorm:"column:last_name" json:"last_name"`
	Email        string  `gorm:"column:email;not null;uniqueIndex" json:"email"`
	PasswordHash string  `gorm:"column:password_hash;not null" json:"-"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
