package main

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rafitanujaya/go-fiber-template/src/database/migrations"
	"github.com/rafitanujaya/go-fiber-template/src/di"
	httpServer "github.com/rafitanujaya/go-fiber-template/src/http"
)

func main() {
	err := godotenv.Load()
	di.HealthCheck()
	if err != nil {
		panic(err)
	}

	//? Auto Migrate
	migrations.Migrate()

	server := httpServer.HttpServer{}
	server.Listen()

}
