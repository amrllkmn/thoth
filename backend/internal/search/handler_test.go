package search

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amrllkmn/thoth/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockSearchService struct {
}

func (m *mockSearchService) FindAll() ([]utils.Book, error) {
	return []utils.Book{
		{Isbn13: 1234567890123, Isbn10: "123457890", Title: "Book 1", Authors: "Author 1"},
		{Isbn13: 9876543210987, Isbn10: "0987654321", Title: "Book 2", Authors: "Author 2"},
	}, nil
}

func (m *mockSearchService) FindByQuery(query string) ([]utils.Book, error) {
	// Mock implementation
	if query == "Book 1" {
		return []utils.Book{
			{Isbn13: 1234567890123, Isbn10: "123457890", Title: "Book 1", Authors: "Author 1"},
		}, nil
	}
	return nil, nil
}
func (m *mockSearchService) FindByID(id uint) {}

func TestHandlerFindAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	handler := NewSQLiteSearchHandler(&mockSearchService{})
	router.GET("/api/v1/books", handler.FindAll)

	req, _ := http.NewRequest("GET", "/api/v1/books", nil)
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
}

func TestHandlerFindByQuery(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	handler := NewSQLiteSearchHandler(&mockSearchService{})
	router.GET("/api/v1/books/search", handler.FindByQuery)

	req, _ := http.NewRequest("GET", "/api/v1/books/search?query=Book 1", nil)
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
