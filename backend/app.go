package main

import (
	"log"
	"os"

	"github.com/amrllkmn/thoth/backend/internal/database"
	"github.com/amrllkmn/thoth/backend/internal/indexed_search"
	"github.com/amrllkmn/thoth/backend/internal/search"
	"github.com/amrllkmn/thoth/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/meilisearch/meilisearch-go"
	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
)

type App struct {
	Router      *gin.Engine
	Db          *gorm.DB
	Meilisearch meilisearch.ServiceManager
}

func setupCORS(router *gin.Engine) {
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "http://localhost:5173"
	}
	config := cors.Config{
		AllowOrigins: []string{allowedOrigins},
		AllowMethods: []string{"GET", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}
	router.Use(cors.New(config))
}

func CreateApp() App {
	router := gin.Default()
	db := database.InitDB()
	meilisearchClient := indexed_search.InitMeilisearchClient()

	setupCORS(router)

	return App{
		Router:      router,
		Db:          db,
		Meilisearch: meilisearchClient,
	}
}

func (app *App) Initialize() {
	err := app.Db.AutoMigrate(&utils.Book{})
	if err != nil {
		log.Fatal("Migration failed")
	}

	sqliteBookRepo := search.NewSQLiteBookRepository(app.Db)
	sqliteSearchService := search.NewSQLiteSearchService(sqliteBookRepo)
	sqliteSearchHandler := search.NewSQLiteSearchHandler(sqliteSearchService)
	search.SetupSQLiteRoutes(app.Router, *sqliteSearchHandler)

	meiliSearchRepo := indexed_search.NewMeilisearchBookRepository(app.Meilisearch)
	meiliSearchService := indexed_search.NewMeilisearchSearchService(meiliSearchRepo)
	meiliSearchHandler := indexed_search.NewMeilisearchHandler(meiliSearchService)
	indexed_search.SetupMeilisearchRoutes(app.Router, *meiliSearchHandler)
}

func (app *App) Run() {
	app.Initialize()
	app.Router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	app.Router.Run(":8080")
}
