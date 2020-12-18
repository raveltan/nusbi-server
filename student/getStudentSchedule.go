package student

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	db "nusbi-server/config"
)

type getStudentScheduleResponse struct{
	Date string
	Course string
	ClassName string
}

func GetStudentScheduleRequest(c *fiber.Ctx) error{
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "s" {
		return c.SendStatus(403)
	}
	id:=c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}
	res, err := db.Db.Query(
		"select s.date_time,c.class_name,cs.course_name from Class c,Enrolled_Courses e,Schedules s,Courses cs where e.student_id = (select student_id from Students where user_id = ?) and e.class_id = c.class_id and s.class_id = c.class_id and cs.course_id = c.course_id",
		id,
	)
	if err != nil {
		return c.SendStatus(500)
	}
	if res.Err() != nil {
		return c.SendStatus(500)
	}
	var result []getStudentScheduleResponse
	for res.Next() {
		var temp getStudentScheduleResponse
		err = res.Scan(&temp.Date, &temp.ClassName, &temp.Course)
		if err != nil {
			return c.SendStatus(500)
		}
		result = append(result, temp)
	}
	return c.JSON(fiber.Map{
		"StudentScheduleResponse": result,
	})
}
