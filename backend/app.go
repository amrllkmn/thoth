package main

import (
	"log"

	"github.com/amrllkmn/thoth/backend/internal/database"
	"github.com/amrllkmn/thoth/backend/internal/search"
	"github.com/amrllkmn/thoth/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Router *gin.Engine
	Db     *gorm.DB
}

func CreateApp() App {
	router := gin.Default()
	db := database.InitDB()

	return App{
		Router: router,
		Db:     db,
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
