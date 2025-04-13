package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" || appEnv == "development" {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file")
		}
	}
	app := CreateApp()
	app.Run()
}
