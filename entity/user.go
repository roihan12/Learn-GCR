package entity

import (
	"time"

	"gorm.io/gorm"
)

// User membuat table/model yang akan dibuat di database
type User struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Name      string         `gorm:"type:varchar(255)" json:"name"`
	Email     string         `gorm:"uniqueIndex; type:varchar(255)" json:"email"`
	Password  string         `gorm:"->;<-;not null" json:"-"`
	Token     string         `gorm:"-" json:"token,omitempty"`
}
