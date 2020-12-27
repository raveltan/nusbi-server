package teacher

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	db "nusbi-server/config"
)

type studentAbsenceRequest struct{
	StudentID string
	ScheduleID string
	Attend bool
}

func SetStudentAbsence(c *fiber.Ctx) error{
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "t" {
		return c.SendStatus(403)
	}
	var req studentAbsenceRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.SendStatus(400)
	}
	if req.Attend {
		_, err = db.Db.Exec(
			"delete from Absence where schedule_id = ? and student_id = ?",
			req.ScheduleID,req.StudentID,
		)
	}else{
		_, err = db.Db.Exec(
			"insert into Absence (schedule_id, student_id) value (?,?)",
			req.ScheduleID,req.StudentID,
		)
	}
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}
