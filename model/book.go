package model

import (
	"time"

	"github.com/gobuffalo/nulls"
)

type Book struct {
	ID          int          `json:"id" gorm:"primary_key"`
	UserID      int          `json:"user_id" gorm:"not null"`
	Title       string       `json:"title" validate:"required" gorm:"not null"`
	Author      string       `json:"author" gorm:"not null"`
	Description nulls.String `json:"description"`
	CreatedAt   time.Time    `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"autoUpdateTime;not null"`
}
