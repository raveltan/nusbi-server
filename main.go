package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"log"
	"nusbi-server/auth"
	"nusbi-server/config"
	"nusbi-server/major"
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
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("yigeiwoligiaogiao"),
	}))
	app.Post("/admin/user/createAdmin", auth.CreateAdmin)
	app.Post("/admin/user/createStudent", auth.CreateStudent)

	app.Post("/admin/major/", major.CreateMajor)
	app.Delete("/admin/major", major.DeleteMajor)
	app.Get("/admin/major", major.GetMajor)

	// Start webserver
	log.Println(app.Listen(config.Port))
}
