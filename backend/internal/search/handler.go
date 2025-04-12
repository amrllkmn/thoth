package search

import (
	"net/http"

	"github.com/amrllkmn/thoth/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type SQLiteSearchHandler struct {
	searchService utils.SearchService
}

func (h *SQLiteSearchHandler) FindAll(c *gin.Context) {
	books, err := h.searchService.FindAll()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"books": books,
		"metadata": gin.H{
			"total": len(books),
		},
	})
}

func (h *SQLiteSearchHandler) FindByQuery(c *gin.Context) {
	query := c.Query("query")
	books, err := h.searchService.FindByQuery(query)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"books": books,
		"metadata": gin.H{
			"total": len(books),
			"query": query,
		},
	})
}

func NewSQLiteSearchHandler(searchService utils.SearchService) *SQLiteSearchHandler {
	return &SQLiteSearchHandler{
		searchService: searchService,
	}
}
