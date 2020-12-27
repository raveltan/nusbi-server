package teacher

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
	db "nusbi-server/config"
)

type getStudentAbsenceScheduleResponse struct {
	Date   string
	ScheduleID string
}

func GetStudentAbsenceSchedule(c *fiber.Ctx) error {
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "t" {
		return c.SendStatus(403)
	}
	id := c.Params("ClassId")
	if id == "" {
		return c.SendStatus(400)
	}
	id2 := c.Params("id")
	if id2 == "" {
		return c.SendStatus(400)
	}
	res, err := db.Db.Query(
		"select s.date_time, s.schedule_id from Schedules s,Class c,Enrolled_Courses e where e.class_id = c.class_id and s.class_id = c.class_id and e.student_id = ? and e.class_id = ?",
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
	var result []getStudentAbsenceScheduleResponse
	for res.Next() {
		var temp getStudentAbsenceScheduleResponse
		err = res.Scan(&temp.Date,&temp.ScheduleID)
		if err != nil {
			log.Println(err)
			return c.SendStatus(500)
		}
		result = append(result, temp)
	}
	return c.JSON(fiber.Map{
		"StudentAbsenceScheduleData": result,
	})
}
