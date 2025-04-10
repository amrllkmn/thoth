package main

import "github.com/gin-gonic/gin"

type App struct {
	Router *gin.Engine
}

func CreateApp() App {
	router := gin.Default()
	return App{
		Router: router,
	}
}

func (app *App) Run() {
	app.Router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	app.Router.Run(":8080")
}
