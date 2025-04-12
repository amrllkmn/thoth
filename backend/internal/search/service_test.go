package search

import (
	"testing"

	"github.com/amrllkmn/thoth/backend/internal/utils"
	"github.com/stretchr/testify/assert"
)

type MockBookRepository struct {
	books []utils.Book
	err   error
}

func (m *MockBookRepository) FindAll(page, limit int) ([]utils.Book, error) {
	return m.books, m.err
}

func (m *MockBookRepository) FindByQuery(query string) ([]utils.Book, error) {
	var filteredBooks []utils.Book
	for _, book := range m.books {
		if book.Title == query || book.Authors == query {
			filteredBooks = append(filteredBooks, book)
		}
	}
	return filteredBooks, m.err
}
func (m *MockBookRepository) FindByID(isbn string) (*utils.Book, error) {
	for _, book := range m.books {
		if book.Isbn13 == isbn || book.Isbn10 == isbn {
			return &book, m.err
		}
	}
	return nil, m.err
}

func setupRepo() *MockBookRepository {
	mockBooks := []utils.Book{
		{Isbn13: "1234567890123", Isbn10: "1234567890", Title: "Book 1", Authors: "Author 1"},
		{Isbn13: "9876543210987", Isbn10: "0987654321", Title: "Book 2", Authors: "Author 2"},
	}

	mockRepo := &MockBookRepository{
		books: mockBooks,
		err:   nil,
	}
	return mockRepo
}

func TestServiceFindAll(t *testing.T) {
	mockRepo := setupRepo()

	service := NewSQLiteSearchService(mockRepo)

	books, err := service.FindAll(0, 0)

	assert.NoError(t, err)
	assert.Len(t, books, 2)

}

func TestServiceFindByQuery(t *testing.T) {
	mockRepo := setupRepo()

	service := NewSQLiteSearchService(mockRepo)

	books, err := service.FindByQuery("Book 1")

	assert.NoError(t, err)
	assert.Len(t, books, 1)
	assert.Equal(t, "Book 1", books[0].Title)
}

func TestServiceFindAll_RepoError(t *testing.T) {
	mockRepo := setupRepo()
	mockRepo.err = assert.AnError
	mockRepo.books = nil
	service := NewSQLiteSearchService(mockRepo)
	books, err := service.FindAll(0, 0)
	assert.Error(t, err)
	assert.Nil(t, books)
}

func TestServiceFindByQuery_RepoError(t *testing.T) {
	mockRepo := setupRepo()
	mockRepo.err = assert.AnError
	mockRepo.books = nil
	service := NewSQLiteSearchService(mockRepo)
	books, err := service.FindByQuery("Book 1")
	assert.Error(t, err)
	assert.Nil(t, books)
}

func TestServiceFindByID(t *testing.T) {
	mockRepo := setupRepo()

	service := NewSQLiteSearchService(mockRepo)

	book, err := service.FindByID("1234567890123")

	assert.NoError(t, err)
	assert.Equal(t, "Book 1", book.Title)
}

func TestServiceFindByID_NotFound(t *testing.T) {
	mockRepo := setupRepo()

	service := NewSQLiteSearchService(mockRepo)

	book, err := service.FindByID("non-existing-isbn")

	assert.NoError(t, err)
	assert.Nil(t, book)
}

func TestServiceFindByID_RepoError(t *testing.T) {
	mockRepo := setupRepo()
	mockRepo.err = assert.AnError
	mockRepo.books = nil
	service := NewSQLiteSearchService(mockRepo)
	book, err := service.FindByID("1234567890123")
	assert.Error(t, err)
	assert.Nil(t, book)
}
