package teacher

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
	db "nusbi-server/config"
)

type getStudentAbsenceResponse struct{
	Absence int
}

func GetStudentAbsence(c*fiber.Ctx) error{
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "t" {
		return c.SendStatus(403)
	}
	id := c.Params("scheduleID")
	if id == "" {
		return c.SendStatus(400)
	}
	id2 := c.Params("id")
	if id2 == "" {
		return c.SendStatus(400)
	}
	res, err := db.Db.Query(
		"select count(*) from Absence where student_id = ? and schedule_id = ?",
		id2, id,
	)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	if res.Err() != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	var result = getStudentAbsenceResponse{Absence: 0}
	for res.Next() {
		err = res.Scan(&result.Absence)
		if err != nil {
			log.Println(err)
			return c.SendStatus(500)
		}
	}
	return c.JSON(result)
}