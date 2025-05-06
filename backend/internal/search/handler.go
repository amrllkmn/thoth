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
		pageInt = 1
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10 // Default limit
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
			"page":  pageInt,
			"limit": limitInt,
		},
	})
}

func (h *SQLiteSearchHandler) FindByQuery(c *gin.Context) {
	var req utils.FindRequest

	err := c.BindJSON(&req)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error,
		})
		return
	}

	if req.LimitInt <= 0 {
		req.LimitInt = 10
	}

	if req.PageInt <= 0 {
		req.PageInt = 1
	}

	books, err := h.searchService.FindByQuery(req.Query, req.PageInt, req.LimitInt)
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
			"query": req.Query,
			"page":  req.PageInt,
			"limit": req.LimitInt,
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
