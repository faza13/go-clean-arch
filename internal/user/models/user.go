package models

import (
	"base/enum"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint64         `json:"id" gorm:"primaryK"`
	Username  string         `json:"username"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Entity    enum.Entity    `json:"entity"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type UserUpdate struct {
	ID       uint64      `json:"id"`
	Username string      `json:"username"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Entity   enum.Entity `json:"entity"`
}

func (UserUpdate) TableName() string {
	return "users"
}
