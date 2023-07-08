package main

import (
	"tulisaja/config"
	"tulisaja/docs"
)

func main() {
	docs.SwaggerInfo.Title = "Tulisaja api documentation"
	docs.SwaggerInfo.Description = "Api documentation for tulisaja app"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := setupRouter(db)
	r.Run()
}
