package recipes

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Name         string
	Description  string
	Ingredients  string
	Instructions string
	Difficult    string
	Time         string
	Serving      string
	Category     string
	User         string
	UserID       uint
	CategoryID   uint
}

type Usecase interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(recipeDomain *Domain) Domain
	Update(id string, recipeDomain *Domain) Domain
	Delete(id string) bool
}

type Repository interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(recipeDomain *Domain) Domain
	Update(id string, recipeDomain *Domain) Domain
	Delete(id string) bool
}
