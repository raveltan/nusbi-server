package courses

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	db "nusbi-server/config"
)

type createCourseRequest struct {
	Name      string
	TeacherID string
	Scu       int
}

type createCourseResponse struct {
	Error int
}

/*
-1: No error
0: Access denies
1: Bodyparsing error
2: Short name
3: invalid scu
*/

func CreateCourse(c *fiber.Ctx) error {
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.JSON(createCourseResponse{Error: 0})
	}
	var request createCourseRequest
	err := c.BodyParser(&request)
	if err != nil {
		return c.JSON(createCourseResponse{Error: 1})
	}
	if len(request.Name) < 5 {
		return c.JSON(createCourseResponse{Error: 2})
	}
	if request.Scu < 0 || request.Scu > 9 {
		return c.JSON(createCourseResponse{Error: 3})
	}
	_, err = db.Db.Exec("insert into Courses (course_id, course_name, lecturer_id, scu) value "+
		"(?,?,?,?)", uuid.New().String(), request.Name, request.TeacherID, request.Scu)
	if err != nil {
		return c.SendStatus(500)
	}
	return c.JSON(createCourseResponse{Error: -1})
}
