package service

import (
	"gorest/filter"
	"gorest/model"
	"gorest/restapi"

	"gorm.io/gorm"
)

// ListBooks lists books.
func ListBooks(db *gorm.DB, user *model.User, pagination *restapi.Pagination, filter *filter.BookFilter) ([]model.Book, int64, error) {
	var (
		books []model.Book
		count int64
	)
	db = db.Scopes(whereUserIs(user), filter.Filter)
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return books, count, db.Scopes(restapi.Paginate(pagination)).Find(&books).Error
}

func CreateBook(db *gorm.DB, book *model.Book) error {
	return db.Create(book).Error
}

func RetrieveBook(db *gorm.DB, id int, user *model.User) (*model.Book, error) {
	var book = &model.Book{
		ID:     id,
		UserID: user.ID,
	}
	return book, db.Scopes(whereUserIs(user)).Find(book).Error
}

func UpdateBook(db *gorm.DB, book *model.Book) error {
	if err := db.Save(book).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBook(db *gorm.DB, book *model.Book) error {
	if err := db.Delete(book).Error; err != nil {
		return err
	}
	return nil
}

func whereUserIs(user *model.User) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		book := model.Book{
			UserID: user.ID,
		}
		return db.Model(&book).Where(&book)
	}
}
