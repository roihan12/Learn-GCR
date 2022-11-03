package entity

import (
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	ID           uint64         `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	Name         string         `gorm:"type:varchar(255)" json:"name"`
	Description  string         `json:"description"`
	Ingredients  string         `json:"ingredients"`
	Instructions string         `json:"instructions"`
	Time         string         `json:"time"`
	Serving      string         `json:"serving"`
	UserID       uint64         `json:"user_id"`
	CategoryID   uint64         `json:"category_id"`
	Category     Category       `json:"category" gorm:"foreignKey:CategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	User         User           `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"publisher"`
}
