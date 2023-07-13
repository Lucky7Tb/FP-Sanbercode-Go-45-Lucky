package main

import (
	"fmt"
	"os"
	"tulisaja/config"
	"tulisaja/docs"

	"github.com/joho/godotenv"
)

func main() {
	docs.SwaggerInfo.Title = "Tulisaja api"
	docs.SwaggerInfo.Description = "Api documentation for tulisaja app"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/api"

	if value, ok := os.LookupEnv("ENVIRONMENT"); ok {
		if value == "DEVELOPMENT" {
			err := godotenv.Load()
			if err != nil {
				panic("Error loading .env file")
			}
		}
	}

	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := setupRouter(db)
	r.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
