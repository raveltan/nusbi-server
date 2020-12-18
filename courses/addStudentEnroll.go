package courses

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	db "nusbi-server/config"
)

type studentEnrollRequest struct {
	Username string
	ClassID  string
}

func AddStudentEnroll(c *fiber.Ctx) error {
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.JSON(createCourseResponse{Error: 0})
	}
	var req studentEnrollRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.SendStatus(400)
	}
	_, err = db.Db.Exec(
		"insert into Enrolled_Courses (enrolled_id, student_id, class_id, mid_score, final_score) value  (?,(select student_id from Students where user_id = ?),?,-1,-1)",
		uuid.New().String(), req.Username, req.ClassID,
	)
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}
