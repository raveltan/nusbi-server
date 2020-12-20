package teacher

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	db "nusbi-server/config"
)

type getTeacherScheduleResponse struct{
	Date string
	Course string
	ClassName string
}

func GetTeacherScheduleRequest(c *fiber.Ctx) error{
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "t" {
		return c.SendStatus(403)
	}
	id:=c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}
	res, err := db.Db.Query(
		"select s.date_time,cs.course_name,c.class_name from Schedules s,Class c,Courses cs,Lecturers l where s.class_id = c.class_id and cs.course_id = c.course_id and l.lecturer_id = cs.lecturer_id and l.user_id = ?",
		id,
	)
	if err != nil {
		return c.SendStatus(500)
	}
	if res.Err() != nil {
		return c.SendStatus(500)
	}
	var result []getTeacherScheduleResponse
	for res.Next() {
		var temp getTeacherScheduleResponse
		err = res.Scan(&temp.Date, &temp.ClassName, &temp.Course)
		if err != nil {
			return c.SendStatus(500)
		}
		result = append(result, temp)
	}
	return c.JSON(fiber.Map{
		"TeacherScheduleResult": result,
	})
}
