package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("thoth.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to init database")
	}
	return db
}
