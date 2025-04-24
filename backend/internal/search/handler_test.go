package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amrllkmn/thoth/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockSearchService struct {
	books []utils.Book
	err   error
}

func setupService() *mockSearchService {
	mockBooks := []utils.Book{
		{Isbn13: "1234567890123", Isbn10: "123457890", Title: "Book 1", Authors: "Author 1"},
		{Isbn13: "9876543210987", Isbn10: "0987654321", Title: "Book 2", Authors: "Author 2"},
	}
	mockService := &mockSearchService{
		books: mockBooks,
		err:   nil,
	}
	return mockService
}

func (m *mockSearchService) FindAll(page, limit int) ([]utils.Book, error) {
	return m.books, m.err
}

func (m *mockSearchService) FindByQuery(query string, page, limit int) ([]utils.Book, error) {
	// Mock implementation
	var filteredBooks []utils.Book
	for _, book := range m.books {
		if book.Title == query || book.Authors == query {
			filteredBooks = append(filteredBooks, book)
		}
	}
	return filteredBooks, m.err
}

func (m *mockSearchService) FindByID(isbn string) (*utils.Book, error) {
	// Mock implementation
	for _, book := range m.books {
		if book.Isbn13 == isbn || book.Isbn10 == isbn {
			return &book, m.err
		}
	}
	return nil, m.err
}

func TestHandlerFindAll(t *testing.T) {

	mockSearchService := setupService()
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	handler := NewSQLiteSearchHandler(mockSearchService)
	router.GET("/api/v1/books/sqlite", handler.FindAll)

	req, _ := http.NewRequest("GET", "/api/v1/books/sqlite", nil)
	resp := httptest.NewRecorder()
	var responseBody map[string]any

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	books, ok := responseBody["books"].([]any)
	assert.True(t, ok)
	assert.Len(t, books, 2)

	metadata, ok := responseBody["metadata"].(map[string]any)
	assert.True(t, ok)

	total_metadata, ok := metadata["total"]
	assert.True(t, ok)
	assert.Equal(t, float64(2), total_metadata)

	page_metadata, ok := metadata["page"]
	assert.True(t, ok)
	assert.Equal(t, float64(1), page_metadata)

	limit_metadata, ok := metadata["limit"]
	assert.True(t, ok)
	assert.Equal(t, float64(10), limit_metadata)
}

func TestHandlerFindAll_Paginated(t *testing.T) {

	mockSearchService := setupService()
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	handler := NewSQLiteSearchHandler(mockSearchService)
	router.GET("/api/v1/books/sqlite", handler.FindAll)

	req, _ := http.NewRequest("GET", "/api/v1/books/sqlite?page=2&limit=30", nil)
	resp := httptest.NewRecorder()
	var responseBody map[string]any

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	books, ok := responseBody["books"].([]any)
	assert.True(t, ok)
	assert.Len(t, books, 2)

	metadata, ok := responseBody["metadata"].(map[string]any)
	assert.True(t, ok)

	total_metadata, ok := metadata["total"]
	assert.True(t, ok)
	assert.Equal(t, float64(2), total_metadata)

	page_metadata, ok := metadata["page"]
	assert.True(t, ok)
	assert.Equal(t, float64(2), page_metadata)

	limit_metadata, ok := metadata["limit"]
	assert.True(t, ok)
	assert.Equal(t, float64(30), limit_metadata)
}

func TestHandlerFindAll_Error(t *testing.T) {
	mockSearchService := setupService()
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	handler := NewSQLiteSearchHandler(mockSearchService)
	router.GET("/api/v1/books/sqlite", handler.FindAll)

	// Simulate an error in the service
	mockSearchService.err = assert.AnError
	mockSearchService.books = nil

	req, _ := http.NewRequest("GET", "/api/v1/books/sqlite", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusInternalServerError, resp.Code)

	fmt.Println(resp.Body.String())

	var responseBody map[string]any
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	message, ok := responseBody["message"].(string)
	assert.True(t, ok)
	assert.Equal(t, "Something went wrong", message)
}

func TestHandlerFindByQuery(t *testing.T) {

	mockSearchService := setupService()
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	handler := NewSQLiteSearchHandler(mockSearchService)
	router.GET("/api/v1/books/sqlite/search", handler.FindByQuery)

	req, _ := http.NewRequest("GET", "/api/v1/books/sqlite/search?query=Book 1", nil)
	resp := httptest.NewRecorder()
	var responseBody map[string]any

	// Assert 200 OK
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	// Assert response body
	books, ok := responseBody["books"].([]any)
	assert.True(t, ok)
	assert.Len(t, books, 1)

	// Assert metadata
	metadata, ok := responseBody["metadata"].(map[string]any)
	assert.True(t, ok)

	total_metadata, ok := metadata["total"]
	assert.True(t, ok)
	assert.Equal(t, float64(1), total_metadata)

	query_metadata, ok := metadata["query"]
	assert.True(t, ok)
	assert.Equal(t, "Book 1", query_metadata)

	// Check if the book title is correct
	book, ok := books[0].(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, "Book 1", book["title"])
}

func TestHandlerFindByQuery_Error(t *testing.T) {
	mockSearchService := setupService()
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	handler := NewSQLiteSearchHandler(mockSearchService)
	router.GET("/api/v1/books/sqlite/search", handler.FindByQuery)

	// Simulate an error in the service
	mockSearchService.err = assert.AnError
	mockSearchService.books = nil

	req, _ := http.NewRequest("GET", "/api/v1/books/sqlite/search?query=Book 1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusInternalServerError, resp.Code)

	fmt.Println(resp.Body.String())

	var responseBody map[string]any
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	message, ok := responseBody["message"].(string)
	assert.True(t, ok)
	assert.Equal(t, "Something went wrong", message)
}

func TestHandlerFindByID(t *testing.T) {
	mockSearchService := setupService()
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	handler := NewSQLiteSearchHandler(mockSearchService)
	router.GET("/api/v1/books/sqlite/:isbn", handler.FindByID)

	req, _ := http.NewRequest("GET", "/api/v1/books/sqlite/1234567890123", nil)
	resp := httptest.NewRecorder()
	var responseBody map[string]any

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	book, ok := responseBody["title"].(string)
	assert.True(t, ok)
	assert.Equal(t, "Book 1", book)
}

func TestHandlerFindByID_NotFound(t *testing.T) {
	mockSearchService := setupService()
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	handler := NewSQLiteSearchHandler(mockSearchService)
	router.GET("/api/v1/books/sqlite/:isbn", handler.FindByID)

	req, _ := http.NewRequest("GET", "/api/v1/books/sqlite/non-existing-isbn", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusNotFound, resp.Code)

	var responseBody map[string]any
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	message, ok := responseBody["message"].(string)
	assert.True(t, ok)
	assert.Equal(t, "Book not found", message)
}

func TestHandlerFindByID_Error(t *testing.T) {
	mockSearchService := setupService()
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	handler := NewSQLiteSearchHandler(mockSearchService)
	router.GET("/api/v1/books/sqlite/:isbn", handler.FindByID)

	// Simulate an error in the service
	mockSearchService.err = assert.AnError
	mockSearchService.books = nil

	req, _ := http.NewRequest("GET", "/api/v1/books/sqlite/1234567890123", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusInternalServerError, resp.Code)

	fmt.Println(resp.Body.String())

	var responseBody map[string]any
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	message, ok := responseBody["message"].(string)
	assert.True(t, ok)
	assert.Equal(t, "Something went wrong", message)
}
