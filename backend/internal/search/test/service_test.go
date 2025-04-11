package test

import (
	"testing"

	"github.com/amrllkmn/thoth/backend/internal/search"
	"github.com/amrllkmn/thoth/backend/internal/utils"
	"github.com/stretchr/testify/assert"
)

type MockBookRepository struct {
	books []utils.Book
	err   error
}

func (m *MockBookRepository) FindAll() ([]utils.Book, error) {
	return m.books, m.err
}

func (m *MockBookRepository) FindByQuery(query string) {}
func (m *MockBookRepository) FindByID(id uint)         {}

func setupTest() *MockBookRepository {
	mockBooks := []utils.Book{
		{Isbn13: 1234567890123, Isbn10: "1234567890", Title: "Book 1", Authors: "Author 1"},
		{Isbn13: 9876543210987, Isbn10: "0987654321", Title: "Book 2", Authors: "Author 2"},
	}

	mockRepo := &MockBookRepository{
		books: mockBooks,
		err:   nil,
	}
	return mockRepo
}

func TestServiceFindAll(t *testing.T) {
	mockRepo := setupTest()

	service := search.NewSQLiteSearchService(mockRepo)

	books, err := service.FindAll()

	assert.NoError(t, err)
	assert.Len(t, books, 2)

}
