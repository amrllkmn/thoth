package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amrllkmn/thoth/backend/internal/search"
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

	handler := search.NewSQLiteSearchHandler(&mockSearchService{})
	router.GET("/api/v1/search", handler.FindAll)

	req, _ := http.NewRequest("GET", "/api/v1/search", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Book 1")
	assert.Contains(t, resp.Body.String(), "Book 2")
}

func TestHandlerFindByQuery(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	handler := search.NewSQLiteSearchHandler(&mockSearchService{})
	router.GET("/api/v1/search", handler.FindByQuery)

	req, _ := http.NewRequest("GET", "/api/v1/search?query=Book 1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Book 1")
	assert.NotContains(t, resp.Body.String(), "Book 2")
}
