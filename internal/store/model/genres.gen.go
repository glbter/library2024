// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGenre = "genres"

// Genre mapped from table <genres>
type Genre struct {
	ID   int16  `gorm:"column:id;type:smallint;primaryKey;autoIncrement:true" json:"id"`
	Name string `gorm:"column:name;type:character varying(64);not null;uniqueIndex" json:"name"`
}

// TableName Genre's table name
func (*Genre) TableName() string {
	return TableNameGenre
}
