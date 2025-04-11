package test

import (
	"testing"

	"github.com/amrllkmn/thoth/backend/internal/search"
	"github.com/amrllkmn/thoth/backend/internal/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&utils.Book{})

	// Seed test data
	db.Create(&utils.Book{Isbn13: 1234567890123, Isbn10: "1234567890", Title: "Book 1", Authors: "Author 1"})
	db.Create(&utils.Book{Isbn13: 9876543210987, Isbn10: "0987654321", Title: "Book 2", Authors: "Author 2"})
	return db
}

func TestRepoFindAll(t *testing.T) {
	db := setupTestDB()
	repo := search.NewSQLiteBookRepository(db)

	books, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, books, 2)
}
