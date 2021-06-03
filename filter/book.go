package filter

import (
	"time"

	"gorm.io/gorm"
)

type BookFilter struct {
	Title        string    `query:"title"`
	Author       string    `query:"author"`
	CreatedAtLte time.Time `query:"created_at_lte"`
	CreatedAtGte time.Time `query:"created_at_gte"`
}

func (f *BookFilter) Filter(db *gorm.DB) *gorm.DB {
	if f.Title != "" {
		db = db.Where("title ILIKE ?", "%"+f.Title+"%")
	}
	if f.Author != "" {
		db = db.Where("author = ?", f.Author)
	}
	if !f.CreatedAtLte.IsZero() {
		db = db.Where("created_at <= ?", f.CreatedAtLte)
	}
	if !f.CreatedAtGte.IsZero() {
		db = db.Where("created_at >= ?", f.CreatedAtGte)
	}
	return db
}
