package main

import (
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rafitanujaya/go-fiber-template/src/database/migrations"
	"github.com/rafitanujaya/go-fiber-template/src/di"
	httpServer "github.com/rafitanujaya/go-fiber-template/src/http"
)

func main() {
	fmt.Printf("Load ENV\n")
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	fmt.Printf("DI Healthcheck\n")
	di.HealthCheck()

	//? Auto Migrate
	fmt.Printf("Migrate\n")
	migrations.Migrate()

	fmt.Printf("Start Server\n")
	server := httpServer.HttpServer{}
	server.Listen()

}
