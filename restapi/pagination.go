package restapi

import "gorm.io/gorm"

type Pagination struct {
	Page     int `query:"page"`
	PageSize int `query:"page_size"`
}

type Page struct {
	Page       int         `json:"page"`
	Count      int64       `json:"count"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}

func NewPagination() *Pagination {
	return &Pagination{
		Page:     1,
		PageSize: 10,
	}
}

func (pagination *Pagination) normalize() {
	if pagination.Page <= 0 {
		pagination.Page = 1
	}
	switch {
	case pagination.PageSize > 100:
		pagination.PageSize = 100
	case pagination.PageSize <= 0:
		pagination.PageSize = 10
	}

}

func NewPageFromPagination(pagination *Pagination, data interface{}, count int64) *Page {
	return &Page{
		Page:       pagination.Page,
		Count:      count,
		TotalPages: (int(count) + pagination.PageSize - 1) / pagination.PageSize,
		Data:       data,
	}
}

func Paginate(pagination *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pagination.normalize()
		offset := (pagination.Page - 1) * pagination.PageSize
		return db.Offset(offset).Limit(pagination.PageSize)
	}
}
