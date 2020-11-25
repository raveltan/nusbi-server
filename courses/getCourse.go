package courses

import "github.com/gofiber/fiber/v2"

type getCourse struct {
	Data []course
}

type course struct {
	CourseID    string
	Name        string
	Scu         int
	TeacherName string
}

func GetCourse(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
