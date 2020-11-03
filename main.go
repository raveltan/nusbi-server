package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"log"
	"nusbi-server/auth"
	"nusbi-server/config"
)

func main() {
	// Initialize config
	err := config.InitDatabase()
	if err != nil {
		panic(err)
	}
	// Configure webserver
	app := fiber.New()
	// No auth routes
	app.Post("/login", auth.Login)

	// Refresh token route

	// Auth routes
	app.Post("/createUser", auth.CreateUser)

	// Start webserver
	log.Println(app.Listen(config.Port))
}
