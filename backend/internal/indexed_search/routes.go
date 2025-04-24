package indexed_search

import "github.com/gin-gonic/gin"

func SetupMeilisearchRoutes(r *gin.Engine, handler MeilisearchHandler) {
	indexedSearchRoutes := r.Group("/v1/books/meilisearch")
	{
		indexedSearchRoutes.GET("/", handler.FindAll)
		indexedSearchRoutes.GET("/search", handler.FindByQuery)
		indexedSearchRoutes.GET("/:isbn", handler.FindByID)
	}
}
