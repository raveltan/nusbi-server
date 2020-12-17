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
	"nusbi-server/courses/class"
	"nusbi-server/courses/schedule"
	"nusbi-server/major"
	"nusbi-server/student"
	"nusbi-server/teacher"
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

	// REmove on production
	app.Post("/createUser", auth.CreateUser)

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

	app.Post("/admin/user/createAdmin", auth.CreateAdmin)
	app.Post("/admin/user/createStudent", auth.CreateStudent)
	app.Post("/admin/user/createTeacher", auth.CreateTeacher)

	app.Get("/admin/user/", auth.GetUserList)

	app.Post("/admin/major/", major.CreateMajor)
	app.Delete("/admin/major", major.DeleteMajor)
	app.Get("/admin/major", major.GetMajor)

	app.Post("/admin/course",courses.CreateCourse)
	app.Get("/admin/course",courses.GetCourse)
	app.Get("/admin/lecturer",courses.GetLecturer)

	app.Post("/admin/class",class.CreateClass)
	app.Get("/admin/class/:id",class.GetClasses)

	app.Post("/admin/schedule",schedule.CreateSchedule)
	app.Get("/admin/schedule/:id",schedule.GetSchedule)
	app.Delete("/admin/schedule/:id",schedule.DeleteSchedule)

	app.Get("/student/profile/:id",student.GetStudentProfile)

	app.Get("/teacher/profile/:id",teacher.GetTeacherProfile)
	// Start webserver
	log.Println(app.Listen(config.Port))
}
