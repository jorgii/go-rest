package service

import (
	"gorest/model"
	"gorest/restapi"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *model.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func RetrieveUserByEmail(db *gorm.DB, email string) (*model.User, error) {
	var user = model.User{
		Email: email,
	}
	return &user, db.Where(user).First(&user).Error
}

func ListUsers(db *gorm.DB, pagination *restapi.Pagination) ([]model.User, int64, error) {
	var (
		users []model.User
		count int64
	)
	if err := db.Model(&users).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return users, count, db.Scopes(restapi.Paginate(pagination)).Find(&users).Error
}
