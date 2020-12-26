package teacher

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	db "nusbi-server/config"
)

type studentMidScoreData struct {
	StudentID string
	ClassID   string
	MidScore  int
}

func SetMidScore(c *fiber.Ctx) error {
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "t" {
		return c.SendStatus(403)
	}
	var req studentMidScoreData
	err := c.BodyParser(&req)
	if err != nil {
		return c.SendStatus(400)
	}
	_, err = db.Db.Exec(
		"update Enrolled_Courses set mid_score = ? where class_id = ? and student_id = ?",
		req.MidScore, req.ClassID, req.StudentID,
	)
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}
