package search

import "github.com/gin-gonic/gin"

func SetupSQLiteRoutes(r *gin.Engine, handler SQLiteSearchHandler) {
	searchRoutes := r.Group("/v1/books/sqlite")
	{
		searchRoutes.GET("/", handler.FindAll)
		searchRoutes.POST("/search", handler.FindByQuery)
		searchRoutes.GET("/:isbn", handler.FindByID)
	}
}
