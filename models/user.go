package models

import (
	"time"
)

const TableNameUser = "users"

// User model
type User struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Email        string    `gorm:"column:email;not null" json:"email"`
	Name         string    `gorm:"column:name;not null" json:"name"`
	Age          int32     `gorm:"column:age" json:"age"`
	Birthday     time.Time `gorm:"column:birthday" json:"birthday"`
	MemberNumber string    `gorm:"column:member_number" json:"member_number"`
	CreatedAt    time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
