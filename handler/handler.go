package handler

import (
	"gorest/config"
	"gorest/model"

	"gorm.io/gorm"
)

type Handler struct {
	Config *config.Config
	DB     *gorm.DB
	User   *model.User
}

func New(db *gorm.DB, cfg *config.Config) *Handler {
	return &Handler{
		Config: cfg,
		DB:     db,
	}
}
