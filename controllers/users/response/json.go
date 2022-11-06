package response

import (
	"echo-recipe/businesses/users"
	"time"

	"gorm.io/gorm"
)

type UserRegister struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromDomainRegister(domain users.Domain) UserRegister {
	return UserRegister{
		ID:        domain.ID,
		Name:      domain.Name,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

type UserLogin struct {
	Message string `json:"message"`
	Email   string `json:"email"`
	Token   string `json:"token"`
}

func FromDomain(domain users.Domain) UserLogin {
	return UserLogin{
		Message: "Login Success",
		Email:   domain.Email,
		Token:   domain.Token,
	}
}
