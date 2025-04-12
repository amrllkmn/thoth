package search

import (
	"testing"

	"github.com/amrllkmn/thoth/backend/internal/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&utils.Book{})

	// Seed test data
	db.Create(&utils.Book{Isbn13: "1234567890123", Isbn10: "1234567890", Title: "Book 1", Authors: "Author 1"})
	db.Create(&utils.Book{Isbn13: "9876543210987", Isbn10: "0987654321", Title: "Book 2", Authors: "Author 2"})
	return db
}

func TestRepoFindAll(t *testing.T) {
	db := setupTestDB()
	repo := NewSQLiteBookRepository(db)

	books, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, books, 2)
}

func TestRepoFindByQuery(t *testing.T) {
	db := setupTestDB()
	repo := NewSQLiteBookRepository(db)

	books, err := repo.FindByQuery("Book 1")
	assert.NoError(t, err)
	assert.Len(t, books, 1)
	assert.Equal(t, "Book 1", books[0].Title)
}

func TestRepoFindAll_DBError(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	repo := NewSQLiteBookRepository(db)

	// Simulate a DB error
	db.Migrator().DropTable(&utils.Book{})

	books, err := repo.FindAll()
	assert.Error(t, err)
	assert.Nil(t, books)
}

func TestRepoFindByQuery_DBError(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	repo := NewSQLiteBookRepository(db)

	// Simulate a DB error
	db.Migrator().DropTable(&utils.Book{})

	books, err := repo.FindByQuery("Book 1")
	assert.Error(t, err)
	assert.Nil(t, books)
}

func TestRepoFindByID_ISBN13(t *testing.T) {
	db := setupTestDB()
	repo := NewSQLiteBookRepository(db)

	book, err := repo.FindByID("1234567890123")
	assert.NoError(t, err)
	assert.Equal(t, "Book 1", book.Title)
}

func TestRepoFindByID_ISBN10(t *testing.T) {
	db := setupTestDB()
	repo := NewSQLiteBookRepository(db)

	book, err := repo.FindByID("1234567890") // ISBN10
	assert.NoError(t, err)
	assert.Equal(t, "Book 1", book.Title)
}

func TestRepoFindByID_NotFound(t *testing.T) {
	db := setupTestDB()
	repo := NewSQLiteBookRepository(db)

	book, err := repo.FindByID("0000000000000") // Non-existing ISBN
	assert.NoError(t, err)
	assert.Nil(t, book)
}

func TestRepoFindByID_DBError(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	repo := NewSQLiteBookRepository(db)

	// Simulate a DB error
	db.Migrator().DropTable(&utils.Book{})

	books, err := repo.FindByID("1234567890123")
	assert.Error(t, err)
	assert.Nil(t, books)
}
