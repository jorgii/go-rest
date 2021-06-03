package model

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	FirstName string    `json:"first_name" gorm:"not null" validate:"required"`
	LastName  string    `json:"last_name" gorm:"not null" validate:"required"`
	Email     string    `json:"email" gorm:"not null;uniqueIndex" validate:"required,email"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:nano;not null"`
	Books     []Book    `json:"-" gorm:"constraint:OnDelete:CASCADE"`
}
