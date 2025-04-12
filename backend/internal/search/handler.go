package search

import (
	"net/http"
	"strconv"

	"github.com/amrllkmn/thoth/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type SQLiteSearchHandler struct {
	searchService utils.SearchService
}

func (h *SQLiteSearchHandler) FindAll(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid page parameter",
		})
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid limit parameter",
		})
		return
	}

	books, err := h.searchService.FindAll(pageInt, limitInt)
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

func (h *SQLiteSearchHandler) FindByID(c *gin.Context) {
	isbn := c.Param("isbn")
	book, err := h.searchService.FindByID(isbn)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}
	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
		return
	}
	c.JSON(http.StatusOK, book)
}

func NewSQLiteSearchHandler(searchService utils.SearchService) *SQLiteSearchHandler {
	return &SQLiteSearchHandler{
		searchService: searchService,
	}
}
