package models

import (
	"time"
)

// User User Model
type User struct {
	ID       uint32 `gorm:"primary_key;column:id" json:"id,omitempty"`
	Email    string `gorm:"column:email;unique" json:"email,omitempty"`
	Name     string `gorm:"column:name" json:"name,omitempty"`
	Username string `gorm:"column:username;unique" json:"username,omitempty" form:"username" binding:"required"`
	Password string `gorm:"column:password" json:"-" form:"password" json:"password" binding:"required"`
	Address  string `gorm:"column:address;type:text" json:"address,omitempty"`

	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}
