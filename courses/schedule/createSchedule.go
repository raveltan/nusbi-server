package schedule

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	db "nusbi-server/config"
)

type scheduleRequest struct{
	Date string
	ClassID string
}

func CreateSchedule(c *fiber.Ctx)error{
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.SendStatus(403)
	}
	var request scheduleRequest

	err := c.BodyParser(&request)
	if err != nil {
		return c.SendStatus(400)
	}
	_, err = db.Db.Exec("insert into Schedules (schedule_id, date_time, class_id) value (?,?,?)",
		uuid.New().String(),request.Date,request.ClassID)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}