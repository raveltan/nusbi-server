package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v2"
	"log"
	"nusbi-server/auth"
	"nusbi-server/config"
	"nusbi-server/courses"
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
	app.Use(cors.New())
	// No auth routes
	app.Post("/login", auth.Login)

	// Refresh token route
	app.Use("/refresh",jwtware.New(jwtware.Config{
		SigningKey: []byte("aoligei"),
	}))
	app.Post("/refresh", auth.RefreshToken)

	// Auth routes
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("yigeiwoligiaogiao"),
	}))

	app.Post("/createUser", auth.CreateUser)


	app.Post("/admin/user/createAdmin", auth.CreateAdmin)
	app.Post("/admin/user/createStudent", auth.CreateStudent)
	app.Post("/admin/user/createTeacher", auth.CreateTeacher)

	app.Get("/admin/user/", auth.GetUserList)

	app.Post("/admin/major/", major.CreateMajor)
	app.Delete("/admin/major", major.DeleteMajor)
	app.Get("/admin/major", major.GetMajor)

	app.Post("/admin/course",courses.CreateCourse)

	// Start webserver
	log.Println(app.Listen(config.Port))
}
