package search

import "github.com/gin-gonic/gin"

func SetupSQLiteRoutes(r *gin.Engine, handler SQLiteSearchHandler) {
	searchRoutes := r.Group("/v1/books")
	{
		searchRoutes.GET("/", handler.FindAll)
		searchRoutes.GET("/search", handler.FindByQuery)
		searchRoutes.GET("/:isbn", handler.FindByID)
	}
}
