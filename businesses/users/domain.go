package users

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string
	Email     string
	Password  string
	Token     string
}

type Usecase interface {
	Register(userDomain *Domain) Domain
	Login(email string, password string) Domain
	Update(domain *Domain) Domain
}

type Repository interface {
	Register(userDomain *Domain) Domain
	GetByEmail(email string, password string) Domain
	Update(domain *Domain) Domain
}
