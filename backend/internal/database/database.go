package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./data/thoth.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to init database")
	}
	return db
}

func Paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 20
	}
	offset := (page - 1) * limit
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}
